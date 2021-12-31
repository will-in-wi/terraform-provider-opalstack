package opalstack

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceAccount() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceAccountRead,
		Schema: map[string]*schema.Schema{
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ready": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"email": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"web_servers": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"imap_servers": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"smtp_servers": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"ips": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"created_at": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceAccountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	r := m.(*requester)

	accountResponse, _, err := r.client.AccountApi.AccountInfo(*r.auth)
	if err != nil {
		return handleSwaggerError(err)
	}

	d.Set("state", accountResponse.State)
	d.Set("ready", accountResponse.Ready)
	d.Set("email", accountResponse.Email)
	d.Set("web_servers", accountResponse.WebServers)
	d.Set("imap_servers", accountResponse.ImapServers)
	d.Set("smtp_servers", accountResponse.SmtpServers)
	d.Set("ips", accountResponse.Ips)
	d.Set("created_at", accountResponse.CreatedAt.Format(time.RFC3339))
	d.SetId(accountResponse.Id)

	return diag.Diagnostics{}
}
