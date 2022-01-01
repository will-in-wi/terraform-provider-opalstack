package opalstack

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourcePsqldb() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourcePsqldbRead,
		Schema:      psqldbsSchema(true),
	}
}

func dataSourcePsqldbRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	r := m.(*requester)

	var diags diag.Diagnostics

	psqldbResponse, _, err := r.client.PsqldbApi.PsqldbRead(*r.auth, d.Get("id").(string))
	if err != nil {
		return handleSwaggerError(err)
	}

	d.Set("state", psqldbResponse.State)
	d.Set("ready", psqldbResponse.Ready)
	d.Set("name", psqldbResponse.Name)
	d.Set("server", psqldbResponse.Server)
	d.Set("charset", psqldbResponse.Charset)
	d.Set("dbusers_readwrite", psqldbResponse.DbusersReadwrite)
	d.Set("dbusers_readonly", psqldbResponse.DbusersReadonly)
	d.SetId(psqldbResponse.Id)

	return diags
}
