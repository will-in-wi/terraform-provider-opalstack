package opalstack

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func mailusersSchema(idRequired bool) map[string]*schema.Schema {
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
		"imap_server": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"procmailrc": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"autoresponder_enable": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"autoresponder_subject": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"autoresponder_message": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"autoresponder_noreply": {
			Type:     schema.TypeSet,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
	}
}
