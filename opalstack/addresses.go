package opalstack

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func addressesSchema(idRequired bool) map[string]*schema.Schema {
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
		"source": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"destinations": {
			Type:     schema.TypeSet,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"forwards": {
			Type:     schema.TypeSet,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
	}
}
