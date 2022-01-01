package opalstack

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceMariausers() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceMariausersRead,
		Schema: map[string]*schema.Schema{
			"mariausers": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Resource{Schema: mariausersSchema(false)},
			},
		},
	}
}

func dataSourceMariausersRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	r := m.(*requester)

	var diags diag.Diagnostics

	mariausersResponse, _, err := r.client.MariauserApi.MariauserList(*r.auth)
	if err != nil {
		return handleSwaggerError(err)
	}

	uuids := make([]string, 0, len(mariausersResponse))
	mariausers := make([]map[string]interface{}, 0, len(mariausersResponse))
	for _, elem := range mariausersResponse {
		uuids = append(uuids, elem.Id)
		mariausers = append(mariausers, map[string]interface{}{
			"id":       elem.Id,
			"state":    elem.State,
			"ready":    elem.Ready,
			"name":     elem.Name,
			"server":   elem.Server,
			"external": elem.External,
		})
	}
	if err := d.Set("mariausers", mariausers); err != nil {
		return diag.Errorf("Unable to assign mariausers: %s", err)
	}

	d.SetId(generateIdFromList(uuids))

	return diags
}
