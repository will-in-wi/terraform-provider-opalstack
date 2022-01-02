package opalstack

import (
	"context"
	"terraform-provider-opalstack/swagger"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceOsvar() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceOsvarCreate,
		ReadContext:   resourceOsvarRead,
		UpdateContext: resourceOsvarUpdate,
		DeleteContext: resourceOsvarDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"content": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"osusers": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"global": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"last_updated": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceOsvarCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	r := m.(*requester)

	create := swagger.OsVarCreate{
		Name:    d.Get("name").(string),
		Content: d.Get("content").(string),
		Osusers: stringSetToStringArray(d.Get("osusers").(*schema.Set)),
		Global:  d.Get("global").(bool),
	}

	osvarResponse, _, err := r.client.OsvarApi.OsvarCreate(*r.auth, []swagger.OsVarCreate{create})
	if err != nil {
		return handleSwaggerError(err)
	}

	d.SetId(osvarResponse[0].Id)

	retryErr := waitForResourceReady(ctx, d, osvarChecker(r, d))
	if retryErr != nil {
		return diag.Errorf("failed with error while waiting for osvar to be created: %s", retryErr)
	}

	resourceOsvarRead(ctx, d, m)

	return diag.Diagnostics{}
}

func resourceOsvarRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	r := m.(*requester)

	osvarResponse, _, err := r.client.OsvarApi.OsvarRead(*r.auth, d.Id())
	if err != nil {
		return handleSwaggerError(err)
	}

	d.Set("name", osvarResponse.Name)
	d.Set("content", osvarResponse.Content)
	d.Set("osusers", stringArrayToStringSet(osvarResponse.Osusers))
	d.Set("global", osvarResponse.Global)

	return diag.Diagnostics{}
}

func resourceOsvarUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	r := m.(*requester)

	if d.HasChangesExcept("last_updated") {
		update := swagger.OsVarUpdate{
			Id:      d.Id(),
			Content: d.Get("content").(string),
			Osusers: stringSetToStringArray(d.Get("osusers").(*schema.Set)),
			Global:  d.Get("global").(bool),
		}

		_, _, err := r.client.OsvarApi.OsvarUpdate(*r.auth, []swagger.OsVarUpdate{update})
		if err != nil {
			return handleSwaggerError(err)
		}

		d.Set("last_updated", time.Now().Format(time.RFC850))

		retryErr := waitForResourceReady(ctx, d, osvarChecker(r, d))
		if retryErr != nil {
			return diag.Errorf("failed with error while waiting for osvar to be updated: %s", retryErr)
		}
	}

	return resourceOsvarRead(ctx, d, m)
}

func resourceOsvarDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	r := m.(*requester)

	_, err := r.client.OsvarApi.OsvarDelete(*r.auth, []swagger.OsVarRead{{Id: d.Id()}})
	if err != nil {
		return handleSwaggerError(err)
	}

	retryErr := waitForResourceDestroyed(ctx, d, osvarChecker(r, d))
	if retryErr != nil {
		return diag.Errorf("failed with error while waiting for osvar to be destroyed: %s", retryErr)
	}

	d.SetId("")

	return diag.Diagnostics{}
}

func osvarChecker(r *requester, d *schema.ResourceData) func() (bool, error) {
	return func() (bool, error) {
		osvarResponse, _, err := r.client.OsvarApi.OsvarRead(*r.auth, d.Id())
		return osvarResponse.Ready, err
	}
}
