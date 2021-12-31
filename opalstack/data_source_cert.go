package opalstack

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceCert() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceCertRead,
		Schema:      certificateSchema(true, true),
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

	populateFromCertResponse(d, certResponse)
	d.SetId(certResponse.Id)

	return diags
}
