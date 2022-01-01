package opalstack

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSites() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceSitesRead,
		Schema: map[string]*schema.Schema{
			"sites": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Resource{Schema: sitesSchema(false)},
			},
		},
	}
}

func dataSourceSitesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	r := m.(*requester)

	var diags diag.Diagnostics

	sitesResponse, _, err := r.client.SiteApi.SiteList(*r.auth)
	if err != nil {
		return handleSwaggerError(err)
	}

	uuids := make([]string, 0, len(sitesResponse))
	sites := make([]map[string]interface{}, 0, len(sitesResponse))
	for _, elem := range sitesResponse {
		uuids = append(uuids, elem.Id)
		sites = append(sites, map[string]interface{}{
			"id":                       elem.Id,
			"state":                    elem.State,
			"ready":                    elem.Ready,
			"name":                     elem.Name,
			"server":                   elem.Server,
			"ip4":                      elem.Ip4,
			"ip6":                      elem.Ip6,
			"disabled":                 elem.Disabled,
			"domains":                  elem.Domains,
			"routes":                   flattenSiteRoutes(elem.Routes),
			"generate_le":              elem.GenerateLe,
			"cert":                     elem.Cert,
			"redirect":                 elem.Redirect,
			"le_http_challenge_tokens": elem.LeHttpChallengeTokens,
		})
	}
	if err := d.Set("sites", sites); err != nil {
		return diag.Errorf("Unable to assign sites: %s", err)
	}

	d.SetId(generateIdFromList(uuids))

	return diags
}
