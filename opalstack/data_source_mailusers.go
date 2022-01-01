package opalstack

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceMailusers() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceMailusersRead,
		Schema: map[string]*schema.Schema{
			"mailusers": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Resource{Schema: mailusersSchema(false)},
			},
		},
	}
}

func dataSourceMailusersRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	r := m.(*requester)

	var diags diag.Diagnostics

	mailusersResponse, _, err := r.client.MailuserApi.MailuserList(*r.auth)
	if err != nil {
		return handleSwaggerError(err)
	}

	uuids := make([]string, 0, len(mailusersResponse))
	mailusers := make([]map[string]interface{}, 0, len(mailusersResponse))
	for _, elem := range mailusersResponse {
		uuids = append(uuids, elem.Id)
		mailusers = append(mailusers, map[string]interface{}{
			"id":                    elem.Id,
			"state":                 elem.State,
			"ready":                 elem.Ready,
			"name":                  elem.Name,
			"imap_server":           elem.ImapServer,
			"procmailrc":            elem.Procmailrc,
			"autoresponder_enable":  elem.AutoresponderEnable,
			"autoresponder_subject": elem.AutoresponderSubject,
			"autoresponder_message": elem.AutoresponderMessage,
			"autoresponder_noreply": elem.AutoresponderNoreply,
		})
	}
	if err := d.Set("mailusers", mailusers); err != nil {
		return diag.Errorf("Unable to assign mailusers: %s", err)
	}

	d.SetId(generateIdFromList(uuids))

	return diags
}
