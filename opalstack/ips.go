package opalstack

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func ipSchema(idRequired bool) map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Type:     schema.TypeString,
			Computed: !idRequired,
			Required: idRequired,
		},
		"ip": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"server": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"type": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"primary": {
			Type:     schema.TypeBool,
			Computed: true,
		},
	}
}
