package opalstack

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func serverSchema(id_required bool) map[string]*schema.Schema {
	s := map[string]*schema.Schema{
		"id": {
			Type:     schema.TypeString,
			Computed: !id_required,
			Required: id_required,
		},
		"hostname": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"type": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}

	return s
}
