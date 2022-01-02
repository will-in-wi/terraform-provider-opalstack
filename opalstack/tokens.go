package opalstack

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func tokensSchema(idRequired bool) map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": { // This is the key. Duplicating for clarity.
			Type:     schema.TypeString,
			Computed: !idRequired,
			Required: idRequired,
		},
		"name": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"key": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
}
