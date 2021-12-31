package opalstack

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceApps() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceAppsRead,
		Schema: map[string]*schema.Schema{
			"apps": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Resource{Schema: appsSchema(false)},
			},
		},
	}
}

func dataSourceAppsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	r := m.(*requester)

	var diags diag.Diagnostics

	appsResponse, _, err := r.client.AppApi.AppList(*r.auth)
	if err != nil {
		return handleSwaggerError(err)
	}

	uuids := make([]string, 0, len(appsResponse))
	apps := make([]map[string]interface{}, 0, len(appsResponse))
	for _, elem := range appsResponse {
		uuids = append(uuids, elem.Id)
		apps = append(apps, map[string]interface{}{
			"id":            elem.Id,
			"state":         elem.State,
			"ready":         elem.Ready,
			"name":          elem.Name,
			"server":        elem.Server,
			"osuser":        elem.Osuser,
			"type":          elem.Type_,
			"port":          elem.Port,
			"installer_url": elem.InstallerUrl,
			"json":          jsonStructToFlatMap(*elem.Json),
		})
	}
	if err := d.Set("apps", apps); err != nil {
		return diag.Errorf("Unable to assign apps: %s", err)
	}

	d.SetId(generateIdFromList(uuids))

	return diags
}
