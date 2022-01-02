package opalstack

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourcePsqldbs() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourcePsqldbsRead,
		Schema: map[string]*schema.Schema{
			"psqldbs": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Resource{Schema: psqldbsSchema(false)},
			},
		},
	}
}

func dataSourcePsqldbsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	r := m.(*requester)

	psqldbsResponse, _, err := r.client.PsqldbApi.PsqldbList(*r.auth)
	if err != nil {
		return handleSwaggerError(err)
	}

	uuids := make([]string, 0, len(psqldbsResponse))
	psqldbs := make([]map[string]interface{}, 0, len(psqldbsResponse))
	for _, elem := range psqldbsResponse {
		uuids = append(uuids, elem.Id)
		psqldbs = append(psqldbs, map[string]interface{}{
			"id":                elem.Id,
			"state":             elem.State,
			"ready":             elem.Ready,
			"name":              elem.Name,
			"server":            elem.Server,
			"charset":           elem.Charset,
			"dbusers_readwrite": elem.DbusersReadwrite,
			"dbusers_readonly":  elem.DbusersReadonly,
		})
	}
	if err := d.Set("psqldbs", psqldbs); err != nil {
		return diag.Errorf("Unable to assign psqldbs: %s", err)
	}

	d.SetId(generateIdFromList(uuids))

	return diag.Diagnostics{}
}
