package opalstack

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceCert() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceCertRead,
		Schema:      certificateSchema(),
	}
}

func dataSourceCertRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	r := m.(*requester)

	var diags diag.Diagnostics

	uuid, ok := d.GetOk("uuid")
	if !ok {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "UUID for Cert is blank",
			Detail:   "UUID for certificate data source is not set",
		})
		return diags
	}
	strUuid := uuid.(string)

	certResponse, _, err := r.client.CertApi.CertRead(*r.auth, strUuid)
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("name", certResponse.Name)
	d.Set("cert", certResponse.Cert)
	d.Set("intermediates", certResponse.Intermediates)
	d.Set("key", certResponse.Key)
	d.Set("exp_date", certResponse.ExpDate.Format(time.RFC3339))
	d.Set("is_opalstack_letsencrypt", certResponse.IsOpalstackLetsencrypt)
	d.Set("is_opalstack_shared_cert", certResponse.IsOpalstackSharedCert)
	d.Set("listed_domains", certResponse.ListedDomains)
	d.SetId(certResponse.Id)

	return diags
}
