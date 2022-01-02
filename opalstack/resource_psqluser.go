package opalstack

import (
	"context"
	"terraform-provider-opalstack/swagger"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
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

	retryErr := waitForResourceReady(ctx, d, psqluserChecker(r, d))
	if retryErr != nil {
		return diag.Errorf("failed with error while waiting for psqluser to be created: %s", retryErr)
	}

	resourcePsqluserRead(ctx, d, m)

	return diag.Diagnostics{}
}

func resourcePsqluserRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	r := m.(*requester)

	psqluserResponse, _, err := r.client.PsqluserApi.PsqluserRead(*r.auth, d.Id())
	if err != nil {
		return handleSwaggerError(err)
	}

	d.Set("server", psqluserResponse.Server)
	d.Set("name", psqluserResponse.Name)
	d.Set("external", psqluserResponse.External)

	return diag.Diagnostics{}
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

		retryErr := waitForResourceReady(ctx, d, psqluserChecker(r, d))
		if retryErr != nil {
			return diag.Errorf("failed with error while waiting for psqluser to be updated: %s", retryErr)
		}
	}

	return resourcePsqluserRead(ctx, d, m)
}

func resourcePsqluserDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	r := m.(*requester)

	_, err := r.client.PsqluserApi.PsqluserDelete(*r.auth, []swagger.PsqlUserRead{{Id: d.Id()}})
	if err != nil {
		return handleSwaggerError(err)
	}

	retryErr := waitForResourceDestroyed(ctx, d, psqluserChecker(r, d))
	if retryErr != nil {
		return diag.Errorf("failed with error while waiting for psqluser to be destroyed: %s", retryErr)
	}

	d.SetId("")

	return diag.Diagnostics{}
}

func psqluserChecker(r *requester, d *schema.ResourceData) func() (bool, error) {
	return func() (bool, error) {
		psqluserResponse, _, err := r.client.PsqluserApi.PsqluserRead(*r.auth, d.Id())
		return psqluserResponse.Ready, err
	}
}
