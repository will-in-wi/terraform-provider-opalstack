package opalstack

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceMariauser() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceMariauserRead,
		Schema:      mariausersSchema(true),
	}
}

func dataSourceMariauserRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	r := m.(*requester)

	mariauserResponse, _, err := r.client.MariauserApi.MariauserRead(*r.auth, d.Get("id").(string))
	if err != nil {
		return handleSwaggerError(err)
	}

	d.Set("state", mariauserResponse.State)
	d.Set("ready", mariauserResponse.Ready)
	d.Set("name", mariauserResponse.Name)
	d.Set("server", mariauserResponse.Server)
	d.Set("external", mariauserResponse.External)
	d.SetId(mariauserResponse.Id)

	return diag.Diagnostics{}
}
