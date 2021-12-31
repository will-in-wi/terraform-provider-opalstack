package opalstack

import (
	"context"
	"terraform-provider-opalstack/swagger"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCert() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceCertCreate,
		ReadContext:   resourceCertRead,
		UpdateContext: resourceCertUpdate,
		DeleteContext: resourceCertDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"cert": {
				Type:     schema.TypeString,
				Required: true,
			},
			"intermediates": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"key": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Not including the documented is_opalstack_letsencrypt since it doesn't seem to function.
		},
	}
}

func resourceCertCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	r := m.(*requester)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	body := []swagger.CertCreate{
		{
			Name:          d.Get("name").(string),
			Cert:          d.Get("cert").(string),
			Intermediates: d.Get("intermediates").(string),
			Key:           d.Get("key").(string),
		},
	}

	certResponse, _, err := r.client.CertApi.CertCreate(*r.auth, body)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(certResponse[0].Id)
	populateFromCertResponse(d, certResponse[0])

	return diags
}

func resourceCertRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	r := m.(*requester)

	var diags diag.Diagnostics

	certResponse, _, err := r.client.CertApi.CertRead(*r.auth, d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	populateFromCertResponse(d, certResponse)

	return diags
}

func resourceCertUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceCertRead(ctx, d, m)
}

func resourceCertDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	return diags
}
