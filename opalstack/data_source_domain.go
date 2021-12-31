package opalstack

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceDomain() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceDomainRead,
		Schema:      domainSchema(true),
	}
}

func dataSourceDomainRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	r := m.(*requester)

	domainResponse, _, err := r.client.DomainApi.DomainRead(*r.auth, d.Get("id").(string))
	if err != nil {
		return handleSwaggerError(err)
	}

	d.Set("state", domainResponse.State)
	d.Set("ready", domainResponse.Ready)
	d.Set("name", domainResponse.Name)
	d.Set("dkim_record", domainResponse.DkimRecord)
	d.Set("is_valid_hostname", domainResponse.IsValidHostname)
	d.SetId(domainResponse.Id)

	return diag.Diagnostics{}
}
