package opalstack

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceMariadb() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceMariadbRead,
		Schema:      mariadbsSchema(true),
	}
}

func dataSourceMariadbRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	r := m.(*requester)

	var diags diag.Diagnostics

	mariadbResponse, _, err := r.client.MariadbApi.MariadbRead(*r.auth, d.Get("id").(string))
	if err != nil {
		return handleSwaggerError(err)
	}

	d.Set("state", mariadbResponse.State)
	d.Set("ready", mariadbResponse.Ready)
	d.Set("name", mariadbResponse.Name)
	d.Set("server", mariadbResponse.Server)
	d.Set("charset", mariadbResponse.Charset)
	d.Set("dbusers_readwrite", mariadbResponse.DbusersReadwrite)
	d.Set("dbusers_readonly", mariadbResponse.DbusersReadonly)
	d.SetId(mariadbResponse.Id)

	return diags
}
