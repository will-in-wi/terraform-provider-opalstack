package opalstack

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceDomains() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceDomainsRead,
		Schema: map[string]*schema.Schema{
			"domains": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Resource{Schema: domainSchema(false)},
			},
		},
	}
}

func dataSourceDomainsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	r := m.(*requester)

	domainsResponse, _, err := r.client.DomainApi.DomainList(*r.auth)
	if err != nil {
		return diag.FromErr(err)
	}

	uuids := make([]string, 0, len(domainsResponse))
	domains := make([]map[string]interface{}, 0, len(domainsResponse))
	for _, elem := range domainsResponse {
		uuids = append(uuids, elem.Id)
		domains = append(domains, map[string]interface{}{
			"id":                elem.Id,
			"state":             elem.State,
			"ready":             elem.Ready,
			"name":              elem.Name,
			"dkim_record":       elem.DkimRecord,
			"is_valid_hostname": elem.IsValidHostname,
		})
	}
	if err := d.Set("domains", domains); err != nil {
		return diag.Errorf("Unable to assign domains: %s", err)
	}

	d.SetId(generateIdFromList(uuids))

	return diag.Diagnostics{}
}
