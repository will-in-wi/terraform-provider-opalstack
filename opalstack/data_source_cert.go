package opalstack

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceCert() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceCertRead,
		Schema: map[string]*schema.Schema{
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"cert": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"intermediates": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"key": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"exp_date": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_opalstack_letsencrypt": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_opalstack_shared_cert": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"listed_domains": &schema.Schema{
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
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
	d.Set("exp_date", certResponse.ExpDate)
	d.Set("is_opalstack_letsencrypt", certResponse.IsOpalstackLetsencrypt)
	d.Set("is_opalstack_shared_cert", certResponse.IsOpalstackSharedCert)
	d.Set("listed_domains", certResponse.ListedDomains)
	d.SetId(certResponse.Id)

	return diags
}
