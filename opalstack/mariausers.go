package opalstack

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func mariausersSchema(idRequired bool) map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Type:     schema.TypeString,
			Computed: !idRequired,
			Required: idRequired,
		},
		"state": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"ready": {
			Type:     schema.TypeBool,
			Computed: true,
		},
		"name": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"server": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"external": {
			Type:     schema.TypeBool,
			Computed: true,
		},
	}
}
