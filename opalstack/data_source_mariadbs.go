package opalstack

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceMariadbs() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceMariadbsRead,
		Schema: map[string]*schema.Schema{
			"mariadbs": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Resource{Schema: mariadbsSchema(false)},
			},
		},
	}
}

func dataSourceMariadbsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	r := m.(*requester)

	mariadbsResponse, _, err := r.client.MariadbApi.MariadbList(*r.auth)
	if err != nil {
		return handleSwaggerError(err)
	}

	uuids := make([]string, 0, len(mariadbsResponse))
	mariadbs := make([]map[string]interface{}, 0, len(mariadbsResponse))
	for _, elem := range mariadbsResponse {
		uuids = append(uuids, elem.Id)
		mariadbs = append(mariadbs, map[string]interface{}{
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
	if err := d.Set("mariadbs", mariadbs); err != nil {
		return diag.Errorf("Unable to assign mariadbs: %s", err)
	}

	d.SetId(generateIdFromList(uuids))

	return diag.Diagnostics{}
}
