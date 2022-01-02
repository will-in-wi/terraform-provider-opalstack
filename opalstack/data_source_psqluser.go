package opalstack

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourcePsqluser() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourcePsqluserRead,
		Schema:      psqlusersSchema(true),
	}
}

func dataSourcePsqluserRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	r := m.(*requester)

	psqluserResponse, _, err := r.client.PsqluserApi.PsqluserRead(*r.auth, d.Get("id").(string))
	if err != nil {
		return handleSwaggerError(err)
	}

	d.Set("state", psqluserResponse.State)
	d.Set("ready", psqluserResponse.Ready)
	d.Set("name", psqluserResponse.Name)
	d.Set("server", psqluserResponse.Server)
	d.Set("external", psqluserResponse.External)
	d.SetId(psqluserResponse.Id)

	return diag.Diagnostics{}
}
