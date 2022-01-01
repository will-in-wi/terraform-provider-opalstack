package opalstack

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourcePsqlusers() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourcePsqlusersRead,
		Schema: map[string]*schema.Schema{
			"psqlusers": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Resource{Schema: psqlusersSchema(false)},
			},
		},
	}
}

func dataSourcePsqlusersRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	r := m.(*requester)

	var diags diag.Diagnostics

	psqlusersResponse, _, err := r.client.PsqluserApi.PsqluserList(*r.auth)
	if err != nil {
		return handleSwaggerError(err)
	}

	uuids := make([]string, 0, len(psqlusersResponse))
	psqlusers := make([]map[string]interface{}, 0, len(psqlusersResponse))
	for _, elem := range psqlusersResponse {
		uuids = append(uuids, elem.Id)
		psqlusers = append(psqlusers, map[string]interface{}{
			"id":       elem.Id,
			"state":    elem.State,
			"ready":    elem.Ready,
			"name":     elem.Name,
			"server":   elem.Server,
			"external": elem.External,
		})
	}
	if err := d.Set("psqlusers", psqlusers); err != nil {
		return diag.Errorf("Unable to assign psqlusers: %s", err)
	}

	d.SetId(generateIdFromList(uuids))

	return diags
}
