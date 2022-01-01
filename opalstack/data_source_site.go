package opalstack

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSite() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceSiteRead,
		Schema:      sitesSchema(true),
	}
}

func dataSourceSiteRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	r := m.(*requester)

	var diags diag.Diagnostics

	siteResponse, _, err := r.client.SiteApi.SiteRead(*r.auth, d.Get("id").(string))
	if err != nil {
		return handleSwaggerError(err)
	}

	d.Set("state", siteResponse.State)
	d.Set("ready", siteResponse.Ready)
	d.Set("name", siteResponse.Name)
	d.Set("server", siteResponse.Server)
	d.Set("ip4", siteResponse.Ip4)
	d.Set("ip6", siteResponse.Ip6)
	d.Set("disabled", siteResponse.Disabled)
	d.Set("domains", siteResponse.Domains)
	d.Set("routes", flattenSiteRoutes(siteResponse.Routes))
	d.Set("generate_le", siteResponse.GenerateLe)
	d.Set("cert", siteResponse.Cert)
	d.Set("redirect", siteResponse.Redirect)
	d.Set("le_http_challenge_tokens", siteResponse.LeHttpChallengeTokens)
	d.SetId(siteResponse.Id)

	return diags
}
