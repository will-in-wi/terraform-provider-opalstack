package opalstack

import (
	"terraform-provider-opalstack/swagger"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func certificateSchema(uuidRequired, keyIncluded bool) map[string]*schema.Schema {
	fields := map[string]*schema.Schema{
		"uuid": {
			Type:     schema.TypeString,
			Required: uuidRequired,
			Computed: !uuidRequired,
		},
		"name": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"cert": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"intermediates": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"exp_date": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"is_opalstack_letsencrypt": {
			Type:     schema.TypeBool,
			Computed: true,
		},
		"is_opalstack_shared_cert": {
			Type:     schema.TypeBool,
			Computed: true,
		},
		"listed_domains": {
			Type:     schema.TypeSet,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
	}

	if keyIncluded {
		fields["key"] = &schema.Schema{
			Type:     schema.TypeString,
			Computed: true,
		}
	}

	return fields
}

func populateFromCertResponse(d *schema.ResourceData, resp swagger.CertResponse) {
	d.Set("name", resp.Name)
	d.Set("cert", resp.Cert)
	d.Set("intermediates", resp.Intermediates)
	d.Set("key", resp.Key)
	d.Set("exp_date", resp.ExpDate.Format(time.RFC3339))
	d.Set("is_opalstack_letsencrypt", resp.IsOpalstackLetsencrypt)
	d.Set("is_opalstack_shared_cert", resp.IsOpalstackSharedCert)
	d.Set("listed_domains", resp.ListedDomains)
}
