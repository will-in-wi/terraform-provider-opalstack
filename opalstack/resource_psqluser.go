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

func resourcePsqluser() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourcePsqluserCreate,
		ReadContext:   resourcePsqluserRead,
		UpdateContext: resourcePsqluserUpdate,
		DeleteContext: resourcePsqluserDelete,
		Schema: map[string]*schema.Schema{
			"server": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"password": {
				Type:      schema.TypeString,
				Optional:  true,
				Sensitive: true,
			},
			"external": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourcePsqluserCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	r := m.(*requester)

	create := swagger.PsqlUserCreate{
		Server:   d.Get("server").(string),
		Name:     d.Get("name").(string),
		Password: d.Get("password").(string),
		External: d.Get("external").(bool),
	}

	psqluserResponse, _, err := r.client.PsqluserApi.PsqluserCreate(*r.auth, []swagger.PsqlUserCreate{create})
	if err != nil {
		return handleSwaggerError(err)
	}

	d.SetId(psqluserResponse[0].Id)

	retryErr := waitForPsqluserReady(ctx, d, r)
	if retryErr != nil {
		return diag.Errorf("failed with error while waiting for psqluser to be created: %s", retryErr)
	}

	resourcePsqluserRead(ctx, d, m)

	return diag.Diagnostics{}
}

func resourcePsqluserRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	r := m.(*requester)

	var diags diag.Diagnostics

	psqluserResponse, _, err := r.client.PsqluserApi.PsqluserRead(*r.auth, d.Id())
	if err != nil {
		return handleSwaggerError(err)
	}

	d.Set("server", psqluserResponse.Server)
	d.Set("name", psqluserResponse.Name)
	d.Set("external", psqluserResponse.External)

	return diags
}

func resourcePsqluserUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	r := m.(*requester)

	if d.HasChangesExcept("last_updated") {
		update := swagger.PsqlUserUpdate{
			Id:       d.Id(),
			Password: d.Get("password").(string),
			External: d.Get("external").(bool),
		}

		_, _, err := r.client.PsqluserApi.PsqluserUpdate(*r.auth, []swagger.PsqlUserUpdate{update})
		if err != nil {
			return handleSwaggerError(err)
		}

		d.Set("last_updated", time.Now().Format(time.RFC850))

		retryErr := waitForPsqluserReady(ctx, d, r)
		if retryErr != nil {
			return diag.Errorf("failed with error while waiting for psqluser to be updated: %s", retryErr)
		}
	}

	return resourcePsqluserRead(ctx, d, m)
}

func resourcePsqluserDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	r := m.(*requester)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	_, err := r.client.PsqluserApi.PsqluserDelete(*r.auth, []swagger.PsqlUserRead{{Id: d.Id()}})
	if err != nil {
		return handleSwaggerError(err)
	}

	d.SetId("")

	return diags
}

func waitForPsqluserReady(ctx context.Context, d *schema.ResourceData, r *requester) error {
	return resource.RetryContext(ctx, d.Timeout(schema.TimeoutCreate)-time.Minute, func() *resource.RetryError {
		var err error
		psqluserResponse, _, err := r.client.PsqluserApi.PsqluserRead(*r.auth, d.Id())
		if err != nil {
			return resource.NonRetryableError(err)
		}

		if !psqluserResponse.Ready {
			return resource.RetryableError(fmt.Errorf("not ready yet"))
		}

		return nil
	})
}
