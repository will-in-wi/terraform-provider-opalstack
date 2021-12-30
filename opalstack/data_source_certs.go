package opalstack

import (
	"context"
	"crypto/sha256"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceCerts() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceCertsRead,
		Schema: map[string]*schema.Schema{
			"certs": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: certificateSchema(),
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
			"uuid":                     elem.Id,
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

	d.SetId(generateId(uuids))
	if err := d.Set("certs", certs); err != nil {
		return diag.Errorf("Unable to assign certs: %s", err)
	}

	return diags
}

func generateId(uuids []string) string {
	joinedString := strings.Join(uuids, "")
	hash := sha256.Sum256([]byte(joinedString))
	return fmt.Sprintf("%x", hash)
}
