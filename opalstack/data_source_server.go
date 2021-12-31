package opalstack

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceServer() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceServerRead,
		Schema:      serverSchema(true),
	}
}

func dataSourceServerRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	r := m.(*requester)

	var diags diag.Diagnostics

	serverResponse, _, err := r.client.ServerApi.ServerRead(*r.auth, d.Get("id").(string))
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("hostname", serverResponse.Hostname)
	d.Set("type", serverResponse.Type_)
	d.SetId(serverResponse.Id)

	return diags
}
