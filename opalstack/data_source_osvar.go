package opalstack

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceOsvar() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceOsvarRead,
		Schema:      osvarsSchema(true),
	}
}

func dataSourceOsvarRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	r := m.(*requester)

	osvarResponse, _, err := r.client.OsvarApi.OsvarRead(*r.auth, d.Get("id").(string))
	if err != nil {
		return handleSwaggerError(err)
	}

	d.Set("state", osvarResponse.State)
	d.Set("ready", osvarResponse.Ready)
	d.Set("name", osvarResponse.Name)
	d.Set("content", osvarResponse.Content)
	d.Set("osusers", osvarResponse.Osusers)
	d.Set("global", osvarResponse.Global)
	d.SetId(osvarResponse.Id)

	return diag.Diagnostics{}
}
