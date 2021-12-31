package opalstack

import (
	"context"
	"fmt"
	"terraform-provider-opalstack/swagger"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDomain() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceDomainCreate,
		ReadContext:   resourceDomainRead,
		DeleteContext: resourceDomainDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"dkim_record": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_valid_hostname": {
				Type:     schema.TypeBool,
				Computed: true,
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceDomainCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	r := m.(*requester)

	body := []swagger.DomainCreate{
		{
			Name: d.Get("name").(string),
		},
	}

	domainResponse, _, err := r.client.DomainApi.DomainCreate(*r.auth, body)
	if err != nil {
		return handleSwaggerError(err)
	}

	d.SetId(domainResponse[0].Id)

	retryErr := waitForDomainReady(ctx, d, r)
	if retryErr != nil {
		return diag.Errorf("failed with error while waiting for user to be updated: %s", retryErr)
	}

	resourceDomainRead(ctx, d, m)

	return diag.Diagnostics{}
}

func resourceDomainRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	r := m.(*requester)

	domainResponse, _, err := r.client.DomainApi.DomainRead(*r.auth, d.Id())
	if err != nil {
		return handleSwaggerError(err)
	}

	d.Set("name", domainResponse.Name)
	d.Set("dkim_record", domainResponse.DkimRecord)
	d.Set("is_valid_hostname", domainResponse.IsValidHostname)

	return diag.Diagnostics{}
}

func resourceDomainDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	r := m.(*requester)

	_, err := r.client.DomainApi.DomainDelete(*r.auth, []swagger.DomainDelete{{Id: d.Id()}})
	if err != nil {
		return handleSwaggerError(err)
	}

	d.SetId("")

	return diag.Diagnostics{}
}

func waitForDomainReady(ctx context.Context, d *schema.ResourceData, r *requester) error {
	return resource.RetryContext(ctx, d.Timeout(schema.TimeoutCreate)-time.Minute, func() *resource.RetryError {
		var err error
		domainResponse, _, err := r.client.DomainApi.DomainRead(*r.auth, d.Id())
		if err != nil {
			return resource.NonRetryableError(err)
		}

		if !domainResponse.Ready {
			return resource.RetryableError(fmt.Errorf("not ready yet"))
		}

		return nil
	})
}
