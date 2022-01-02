package opalstack

import (
	"context"
	"terraform-provider-opalstack/swagger"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func allowedDnsrecordTypes() []string {
	return []string{
		// See swagger/model_dns_record_type_enum.go for list.
		"A",
		"AAAA",
		"CNAME",
		"MX",
		"TXT",
		"SRV",
	}
}

func resourceDnsrecord() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceDnsrecordCreate,
		ReadContext:   resourceDnsrecordRead,
		UpdateContext: resourceDnsrecordUpdate,
		DeleteContext: resourceDnsrecordDelete,
		Schema: map[string]*schema.Schema{
			"domain": {
				Type:     schema.TypeString,
				Required: true,
			},
			"type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice(allowedDnsrecordTypes(), false),
				Required:     true,
			},
			"content": {
				Type:     schema.TypeString,
				Required: true,
			},
			"priority": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10,
			},
			"ttl": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  3600,
			},
			"last_updated": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceDnsrecordCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	r := m.(*requester)

	dnsrecordType := swagger.DnsRecordTypeEnum(d.Get("type").(string))

	create := swagger.DnsRecordCreate{
		Domain:   d.Get("domain").(string),
		Type_:    &dnsrecordType,
		Content:  d.Get("content").(string),
		Priority: int32(d.Get("priority").(int)),
		Ttl:      int32(d.Get("ttl").(int)),
	}

	dnsrecordResponse, _, err := r.client.DnsrecordApi.DnsrecordCreate(*r.auth, []swagger.DnsRecordCreate{create})
	if err != nil {
		return handleSwaggerError(err)
	}

	d.SetId(dnsrecordResponse[0].Id)

	retryErr := waitForResourceReady(ctx, d, dnsrecordChecker(r, d))
	if retryErr != nil {
		return diag.Errorf("failed with error while waiting for user to be updated: %s", retryErr)
	}

	resourceDnsrecordRead(ctx, d, m)

	return diag.Diagnostics{}
}

func resourceDnsrecordRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	r := m.(*requester)

	dnsrecordResponse, _, err := r.client.DnsrecordApi.DnsrecordRead(*r.auth, d.Id())
	if err != nil {
		return handleSwaggerError(err)
	}

	d.Set("domain", dnsrecordResponse.Domain)
	d.Set("type", dnsrecordResponse.Type_)
	d.Set("content", dnsrecordResponse.Content)
	d.Set("priority", dnsrecordResponse.Priority)
	d.Set("ttl", dnsrecordResponse.Ttl)

	return diag.Diagnostics{}
}

func resourceDnsrecordUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	r := m.(*requester)

	if d.HasChangesExcept("last_updated") {
		dnsrecordType := swagger.DnsRecordTypeEnum(d.Get("type").(string))
		update := swagger.DnsRecordUpdate{
			Id:       d.Id(),
			Domain:   d.Get("domain").(string),
			Type_:    &dnsrecordType,
			Content:  d.Get("content").(string),
			Priority: int32(d.Get("priority").(int)),
			Ttl:      int32(d.Get("ttl").(int)),
		}

		_, _, err := r.client.DnsrecordApi.DnsrecordUpdate(*r.auth, []swagger.DnsRecordUpdate{update})
		if err != nil {
			return handleSwaggerError(err)
		}

		d.Set("last_updated", time.Now().Format(time.RFC850))

		retryErr := waitForResourceReady(ctx, d, dnsrecordChecker(r, d))
		if retryErr != nil {
			return diag.Errorf("failed with error while waiting for user to be updated: %s", retryErr)
		}
	}

	return resourceDnsrecordRead(ctx, d, m)
}

func resourceDnsrecordDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	r := m.(*requester)

	_, err := r.client.DnsrecordApi.DnsrecordDelete(*r.auth, []swagger.DnsRecordRead{{Id: d.Id()}})
	if err != nil {
		return handleSwaggerError(err)
	}

	retryErr := waitForResourceDestroyed(ctx, d, dnsrecordChecker(r, d))
	if retryErr != nil {
		return diag.Errorf("failed with error while waiting for mariauser to be destroyed: %s", retryErr)
	}

	d.SetId("")

	return diag.Diagnostics{}
}

func dnsrecordChecker(r *requester, d *schema.ResourceData) func() (bool, error) {
	return func() (bool, error) {
		dnsrecordResponse, _, err := r.client.DnsrecordApi.DnsrecordRead(*r.auth, d.Id())
		return dnsrecordResponse.Ready, err
	}
}
