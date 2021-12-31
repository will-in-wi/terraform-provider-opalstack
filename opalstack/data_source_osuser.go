package opalstack

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceOsuser() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceOsuserRead,
		Schema:      osusersSchema(true),
	}
}

func dataSourceOsuserRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	r := m.(*requester)

	var diags diag.Diagnostics

	osuserResponse, _, err := r.client.OsuserApi.OsuserRead(*r.auth, d.Get("id").(string))
	if err != nil {
		return handleSwaggerError(err)
	}

	d.Set("state", osuserResponse.State)
	d.Set("ready", osuserResponse.Ready)
	d.Set("name", osuserResponse.Name)
	d.Set("server", osuserResponse.Server)
	d.SetId(osuserResponse.Id)

	return diags
}
