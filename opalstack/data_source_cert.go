package opalstack

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceCert() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceCertRead,
		Schema:      certificateSchema(true),
	}
}

func dataSourceCertRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	r := m.(*requester)

	certResponse, _, err := r.client.CertApi.CertRead(*r.auth, d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	populateFromCertResponse(d, certResponse)
	d.SetId(certResponse.Id)

	return diag.Diagnostics{}
}
