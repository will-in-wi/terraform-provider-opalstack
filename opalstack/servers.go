package opalstack

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func serverSchema(with_id bool) map[string]*schema.Schema {
	s := map[string]*schema.Schema{
		"hostname": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"type": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}

	if with_id {
		s["id"] = &schema.Schema{
			Type:     schema.TypeString,
			Computed: true,
		}
	}

	return s
}
