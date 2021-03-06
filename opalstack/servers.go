package opalstack

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func serverSchema(idRequired bool) map[string]*schema.Schema {
	s := map[string]*schema.Schema{
		"id": {
			Type:     schema.TypeString,
			Computed: !idRequired,
			Required: idRequired,
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
