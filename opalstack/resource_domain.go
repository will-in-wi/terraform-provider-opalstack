package opalstack

import (
	"context"
	"terraform-provider-opalstack/swagger"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
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
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ready": {
				Type:     schema.TypeBool,
				Computed: true,
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
		return diag.FromErr(err)
	}

	d.SetId(domainResponse[0].Id)
	resourceDomainRead(ctx, d, m)

	return diag.Diagnostics{}
}

func resourceDomainRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	r := m.(*requester)

	domainResponse, _, err := r.client.DomainApi.DomainRead(*r.auth, d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("name", domainResponse.Name)
	d.Set("state", domainResponse.State)
	d.Set("ready", domainResponse.Ready)
	d.Set("dkim_record", domainResponse.DkimRecord)
	d.Set("is_valid_hostname", domainResponse.IsValidHostname)

	return diag.Diagnostics{}
}

func resourceDomainDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	r := m.(*requester)

	_, err := r.client.DomainApi.DomainDelete(*r.auth, []swagger.DomainDelete{{Id: d.Id()}})
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return diag.Diagnostics{}
}
