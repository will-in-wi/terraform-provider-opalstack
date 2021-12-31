package opalstack

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceDnsrecords() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceDnsrecordsRead,
		Schema: map[string]*schema.Schema{
			"dnsrecords": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Resource{Schema: dnsrecordSchema(false)},
			},
		},
	}
}

func dataSourceDnsrecordsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	r := m.(*requester)

	var diags diag.Diagnostics

	dnsrecordsResponse, _, err := r.client.DnsrecordApi.DnsrecordList(*r.auth)
	if err != nil {
		return diag.FromErr(err)
	}

	uuids := make([]string, 0, len(dnsrecordsResponse))
	dnsrecords := make([]map[string]interface{}, 0, len(dnsrecordsResponse))
	for _, elem := range dnsrecordsResponse {
		uuids = append(uuids, elem.Id)
		dnsrecords = append(dnsrecords, map[string]interface{}{
			"id":       elem.Id,
			"state":    elem.State,
			"ready":    elem.Ready,
			"domain":   elem.Domain,
			"type":     elem.Type_,
			"content":  elem.Content,
			"priority": elem.Priority,
			"ttl":      elem.Ttl,
		})
	}
	if err := d.Set("dnsrecords", dnsrecords); err != nil {
		return diag.Errorf("Unable to assign dnsrecords: %s", err)
	}

	d.SetId(generateIdFromList(uuids))

	return diags
}
