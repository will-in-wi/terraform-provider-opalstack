package opalstack

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceCertShared() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceCertSharedRead,
		Schema:      certificateSchema(false, false),
	}
}

func dataSourceCertSharedRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	r := m.(*requester)

	var diags diag.Diagnostics

	certResponse, _, err := r.client.CertApi.CertShared(*r.auth)
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("uuid", certResponse.Id)
	d.Set("name", certResponse.Name)
	d.Set("cert", certResponse.Cert)
	d.Set("intermediates", certResponse.Intermediates)
	d.Set("exp_date", certResponse.ExpDate.Format(time.RFC3339))
	d.Set("is_opalstack_letsencrypt", certResponse.IsOpalstackLetsencrypt)
	d.Set("is_opalstack_shared_cert", certResponse.IsOpalstackSharedCert)
	d.Set("listed_domains", certResponse.ListedDomains)
	d.SetId(certResponse.Id)

	return diags
}
