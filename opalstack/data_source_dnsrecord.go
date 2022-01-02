package opalstack

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceDnsrecord() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceDnsrecordRead,
		Schema:      dnsrecordSchema(true),
	}
}

func dataSourceDnsrecordRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	r := m.(*requester)

	dnsrecordResponse, _, err := r.client.DnsrecordApi.DnsrecordRead(*r.auth, d.Get("id").(string))
	if err != nil {
		return handleSwaggerError(err)
	}

	d.Set("id", dnsrecordResponse.Id)
	d.Set("state", dnsrecordResponse.State)
	d.Set("ready", dnsrecordResponse.Ready)
	d.Set("domain", dnsrecordResponse.Domain)
	d.Set("type", dnsrecordResponse.Type_)
	d.Set("content", dnsrecordResponse.Content)
	d.Set("priority", dnsrecordResponse.Priority)
	d.Set("ttl", dnsrecordResponse.Ttl)
	d.SetId(dnsrecordResponse.Id)

	return diag.Diagnostics{}
}
