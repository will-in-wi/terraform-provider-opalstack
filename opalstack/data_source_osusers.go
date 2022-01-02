package opalstack

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceOsusers() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceOsusersRead,
		Schema: map[string]*schema.Schema{
			"osusers": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Resource{Schema: osusersSchema(false)},
			},
		},
	}
}

func dataSourceOsusersRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	r := m.(*requester)

	osusersResponse, _, err := r.client.OsuserApi.OsuserList(*r.auth)
	if err != nil {
		return handleSwaggerError(err)
	}

	uuids := make([]string, 0, len(osusersResponse))
	osusers := make([]map[string]interface{}, 0, len(osusersResponse))
	for _, elem := range osusersResponse {
		uuids = append(uuids, elem.Id)
		osusers = append(osusers, map[string]interface{}{
			"id":     elem.Id,
			"state":  elem.State,
			"ready":  elem.Ready,
			"name":   elem.Name,
			"server": elem.Server,
		})
	}
	if err := d.Set("osusers", osusers); err != nil {
		return diag.Errorf("Unable to assign osusers: %s", err)
	}

	d.SetId(generateIdFromList(uuids))

	return diag.Diagnostics{}
}
