package opalstack

import (
	"context"
	"terraform-provider-opalstack/swagger"
	"time"

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
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: compareTrimmed,
			},
			"intermediates": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: compareTrimmed,
			},
			"key": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: compareTrimmed,
			},
			"last_updated": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			// Not including the documented is_opalstack_letsencrypt since it doesn't seem to function.
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceCertCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	r := m.(*requester)

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
		return handleSwaggerError(err)
	}

	d.SetId(certResponse[0].Id)
	resourceCertRead(ctx, d, m)

	return diag.Diagnostics{}
}

func resourceCertRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	r := m.(*requester)

	certResponse, _, err := r.client.CertApi.CertRead(*r.auth, d.Id())
	if err != nil {
		return handleSwaggerError(err)
	}

	populateFromCertResponse(d, certResponse)

	return diag.Diagnostics{}
}

func resourceCertUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	r := m.(*requester)

	if d.HasChangesExcept("last_updated") {
		body := []swagger.CertUpdate{
			{
				Id:            d.Id(),
				Name:          d.Get("name").(string),
				Cert:          d.Get("cert").(string),
				Intermediates: d.Get("intermediates").(string),
				Key:           d.Get("key").(string),
			},
		}

		_, _, err := r.client.CertApi.CertUpdate(*r.auth, body)
		if err != nil {
			return handleSwaggerError(err)
		}

		d.Set("last_updated", time.Now().Format(time.RFC850))
	}

	return resourceCertRead(ctx, d, m)
}

func resourceCertDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	r := m.(*requester)

	_, err := r.client.CertApi.CertDelete(*r.auth, []swagger.CertRead{{Id: d.Id()}})
	if err != nil {
		return handleSwaggerError(err)
	}

	d.SetId("")

	return diag.Diagnostics{}
}
