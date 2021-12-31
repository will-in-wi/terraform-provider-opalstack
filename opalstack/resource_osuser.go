package opalstack

import (
	"context"
	"terraform-provider-opalstack/swagger"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceOsuser() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceOsuserCreate,
		ReadContext:   resourceOsuserRead,
		UpdateContext: resourceOsuserUpdate,
		DeleteContext: resourceOsuserDelete,
		Schema: map[string]*schema.Schema{
			"server": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"password": {
				Type:     schema.TypeString,
				Optional: true,
				// TODO: Make this secret.
			},
			"ready": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
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

func resourceOsuserCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	r := m.(*requester)

	create := swagger.OsUserCreate{
		Server:   d.Get("server").(string),
		Name:     d.Get("name").(string),
		Password: d.Get("password").(string),
	}

	osuserResponse, _, err := r.client.OsuserApi.OsuserCreate(*r.auth, []swagger.OsUserCreate{create})
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(osuserResponse[0].Id)
	resourceOsuserRead(ctx, d, m)

	return diag.Diagnostics{}
}

func resourceOsuserRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	r := m.(*requester)

	var diags diag.Diagnostics

	osuserResponse, _, err := r.client.OsuserApi.OsuserRead(*r.auth, d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("server", osuserResponse.Server)
	d.Set("name", osuserResponse.Name)
	d.Set("ready", osuserResponse.Ready)
	d.Set("state", osuserResponse.State)

	return diags
}

func resourceOsuserUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	r := m.(*requester)

	if d.HasChangesExcept("last_updated") {
		update := swagger.OsUserUpdate{
			Id:       d.Id(),
			Password: d.Get("password").(string),
		}

		_, _, err := r.client.OsuserApi.OsuserUpdate(*r.auth, []swagger.OsUserUpdate{update})
		if err != nil {
			return diag.FromErr(err)
		}

		d.Set("last_updated", time.Now().Format(time.RFC850))
	}

	return resourceOsuserRead(ctx, d, m)
}

func resourceOsuserDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	r := m.(*requester)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	_, err := r.client.OsuserApi.OsuserDelete(*r.auth, []swagger.OsUserRead{{Id: d.Id()}})
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return diags
}
