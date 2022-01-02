package opalstack

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceIps() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIpsRead,
		Schema: map[string]*schema.Schema{
			"ips": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Resource{Schema: ipSchema(false)},
			},
		},
	}
}

func dataSourceIpsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	r := m.(*requester)

	ipsResponse, _, err := r.client.IpApi.IpList(*r.auth)
	if err != nil {
		return handleSwaggerError(err)
	}

	uuids := make([]string, 0, len(ipsResponse))
	ips := make([]map[string]interface{}, 0, len(ipsResponse))
	for _, elem := range ipsResponse {
		uuids = append(uuids, elem.Id)
		ips = append(ips, map[string]interface{}{
			"id":      elem.Id,
			"ip":      elem.Ip,
			"server":  elem.Server,
			"type":    elem.Type_,
			"primary": elem.Primary,
		})
	}
	if err := d.Set("ips", ips); err != nil {
		return diag.Errorf("Unable to assign ips: %s", err)
	}

	d.SetId(generateIdFromList(uuids))

	return diag.Diagnostics{}
}
