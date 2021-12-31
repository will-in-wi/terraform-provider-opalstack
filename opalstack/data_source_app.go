package opalstack

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceApp() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceAppRead,
		Schema:      appsSchema(true),
	}
}

func dataSourceAppRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	r := m.(*requester)

	var diags diag.Diagnostics

	appResponse, _, err := r.client.AppApi.AppRead(*r.auth, d.Get("id").(string))
	if err != nil {
		return handleSwaggerError(err)
	}

	d.Set("state", appResponse.State)
	d.Set("ready", appResponse.Ready)
	d.Set("name", appResponse.Name)
	d.Set("server", appResponse.Server)
	d.Set("osuser", appResponse.Osuser)
	d.Set("type", appResponse.Type_)
	d.Set("port", appResponse.Port)
	d.Set("installer_url", appResponse.InstallerUrl)
	d.Set("json", jsonStructToFlatMap(*appResponse.Json))
	d.SetId(appResponse.Id)

	return diags
}
