package opalstack

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceMailuser() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceMailuserRead,
		Schema:      mailusersSchema(true),
	}
}

func dataSourceMailuserRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	r := m.(*requester)

	var diags diag.Diagnostics

	mailuserResponse, _, err := r.client.MailuserApi.MailuserRead(*r.auth, d.Get("id").(string))
	if err != nil {
		return handleSwaggerError(err)
	}

	d.Set("state", mailuserResponse.State)
	d.Set("ready", mailuserResponse.Ready)
	d.Set("name", mailuserResponse.Name)
	d.Set("imap_server", mailuserResponse.ImapServer)
	d.Set("procmailrc", mailuserResponse.Procmailrc)
	d.Set("autoresponder_enable", mailuserResponse.AutoresponderEnable)
	d.Set("autoresponder_subject", mailuserResponse.AutoresponderSubject)
	d.Set("autoresponder_message", mailuserResponse.AutoresponderMessage)
	d.Set("autoresponder_noreply", mailuserResponse.AutoresponderNoreply)
	d.SetId(mailuserResponse.Id)

	return diags
}
