package opalstack

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceIp() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIpRead,
		Schema:      ipSchema(true),
	}
}

func dataSourceIpRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	r := m.(*requester)

	ipResponse, _, err := r.client.IpApi.IpRead(*r.auth, d.Get("id").(string))
	if err != nil {
		return handleSwaggerError(err)
	}

	d.Set("ip", ipResponse.Ip)
	d.Set("server", ipResponse.Server)
	d.Set("type", ipResponse.Type_)
	d.Set("primary", ipResponse.Primary)
	d.SetId(ipResponse.Id)

	return diag.Diagnostics{}
}
