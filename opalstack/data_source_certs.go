package opalstack

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceCerts() *schema.Resource {
	certSchema := certificateSchema(true)
	certSchema["id"] = &schema.Schema{
		Type:     schema.TypeString,
		Computed: true,
	}

	return &schema.Resource{
		ReadContext: dataSourceCertsRead,
		Schema: map[string]*schema.Schema{
			"certs": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: certSchema,
				},
			},
		},
	}
}

func dataSourceCertsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	r := m.(*requester)

	var diags diag.Diagnostics

	certResponse, _, err := r.client.CertApi.CertList(*r.auth)
	if err != nil {
		return diag.FromErr(err)
	}

	certs := make([]map[string]interface{}, 0, len(certResponse))
	uuids := make([]string, 0, len(certResponse))
	for _, elem := range certResponse {
		uuids = append(uuids, elem.Id)
		certs = append(certs, map[string]interface{}{
			"id":                       elem.Id,
			"name":                     elem.Name,
			"cert":                     elem.Cert,
			"intermediates":            elem.Intermediates,
			"key":                      elem.Key,
			"exp_date":                 elem.ExpDate.Format(time.RFC3339),
			"is_opalstack_letsencrypt": elem.IsOpalstackLetsencrypt,
			"is_opalstack_shared_cert": elem.IsOpalstackSharedCert,
			"listed_domains":           elem.ListedDomains,
		})
	}

	d.SetId(generateIdFromList(uuids))
	if err := d.Set("certs", certs); err != nil {
		return diag.Errorf("Unable to assign certs: %s", err)
	}

	return diags
}
