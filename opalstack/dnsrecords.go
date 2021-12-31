package opalstack

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func dnsrecordSchema(idRequired bool) map[string]*schema.Schema {
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
		"domain": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"type": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"content": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"priority": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"ttl": {
			Type:     schema.TypeInt,
			Computed: true,
		},
	}
}
