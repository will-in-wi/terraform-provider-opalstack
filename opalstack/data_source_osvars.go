package opalstack

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceOsvars() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceOsvarsRead,
		Schema: map[string]*schema.Schema{
			"osvars": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Resource{Schema: osvarsSchema(false)},
			},
		},
	}
}

func dataSourceOsvarsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	r := m.(*requester)

	osvarsResponse, _, err := r.client.OsvarApi.OsvarList(*r.auth)
	if err != nil {
		return handleSwaggerError(err)
	}

	uuids := make([]string, 0, len(osvarsResponse))
	osvars := make([]map[string]interface{}, 0, len(osvarsResponse))
	for _, elem := range osvarsResponse {
		uuids = append(uuids, elem.Id)
		osvars = append(osvars, map[string]interface{}{
			"id":      elem.Id,
			"state":   elem.State,
			"ready":   elem.Ready,
			"name":    elem.Name,
			"content": elem.Content,
			"osusers": elem.Osusers,
			"global":  elem.Global,
		})
	}
	if err := d.Set("osvars", osvars); err != nil {
		return diag.Errorf("Unable to assign osvars: %s", err)
	}

	d.SetId(generateIdFromList(uuids))

	return diag.Diagnostics{}
}
