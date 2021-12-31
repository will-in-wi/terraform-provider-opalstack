package opalstack

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceServers() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceServersRead,
		Schema: map[string]*schema.Schema{
			"web_servers": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Resource{Schema: serverSchema(false)},
			},
			"imap_servers": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Resource{Schema: serverSchema(false)},
			},
			"smtp_servers": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Resource{Schema: serverSchema(false)},
			},
		},
	}
}

func dataSourceServersRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	r := m.(*requester)

	var diags diag.Diagnostics

	serversResponse, _, err := r.client.ServerApi.ServerList(*r.auth)
	if err != nil {
		return diag.FromErr(err)
	}

	uuids := make([]string, 0, len(serversResponse.ImapServers)+len(serversResponse.SmtpServers)+len(serversResponse.WebServers))

	imapServers := make([]map[string]interface{}, 0, len(serversResponse.ImapServers))
	for _, elem := range serversResponse.ImapServers {
		uuids = append(uuids, elem.Id)
		imapServers = append(imapServers, map[string]interface{}{
			"id":       elem.Id,
			"hostname": elem.Hostname,
			"type":     elem.Type_,
		})
	}
	if err := d.Set("imap_servers", imapServers); err != nil {
		return diag.Errorf("Unable to assign imap servers: %s", err)
	}

	smtpServers := make([]map[string]interface{}, 0, len(serversResponse.SmtpServers))
	for _, elem := range serversResponse.SmtpServers {
		uuids = append(uuids, elem.Id)
		smtpServers = append(smtpServers, map[string]interface{}{
			"id":       elem.Id,
			"hostname": elem.Hostname,
			"type":     elem.Type_,
		})
	}
	if err := d.Set("smtp_servers", smtpServers); err != nil {
		return diag.Errorf("Unable to assign smtp servers: %s", err)
	}

	webServers := make([]map[string]interface{}, 0, len(serversResponse.WebServers))
	for _, elem := range serversResponse.WebServers {
		uuids = append(uuids, elem.Id)
		webServers = append(webServers, map[string]interface{}{
			"id":       elem.Id,
			"hostname": elem.Hostname,
			"type":     elem.Type_,
		})
	}
	if err := d.Set("web_servers", webServers); err != nil {
		return diag.Errorf("Unable to assign web servers: %s", err)
	}

	d.SetId(generateIdFromList(uuids))

	return diags
}
