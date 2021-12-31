package opalstack

import (
	"context"
	"fmt"
	"terraform-provider-opalstack/swagger"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
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
				ValidateFunc: validateStringInList(allowedDnsrecordTypes()),
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
		return diag.FromErr(err)
	}

	d.SetId(dnsrecordResponse[0].Id)

	retryErr := waitForDnsrecordReady(ctx, d, r)
	if retryErr != nil {
		return diag.Errorf("failed with error while waiting for user to be updated: %s", retryErr)
	}

	resourceDnsrecordRead(ctx, d, m)

	return diag.Diagnostics{}
}

func resourceDnsrecordRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	r := m.(*requester)

	var diags diag.Diagnostics

	dnsrecordResponse, _, err := r.client.DnsrecordApi.DnsrecordRead(*r.auth, d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("domain", dnsrecordResponse.Domain)
	d.Set("type", dnsrecordResponse.Type_)
	d.Set("content", dnsrecordResponse.Content)
	d.Set("priority", dnsrecordResponse.Priority)
	d.Set("ttl", dnsrecordResponse.Ttl)

	return diags
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
			return diag.FromErr(err)
		}

		d.Set("last_updated", time.Now().Format(time.RFC850))

		retryErr := waitForDnsrecordReady(ctx, d, r)
		if retryErr != nil {
			return diag.Errorf("failed with error while waiting for user to be updated: %s", retryErr)
		}
	}

	return resourceDnsrecordRead(ctx, d, m)
}

func resourceDnsrecordDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	r := m.(*requester)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	_, err := r.client.DnsrecordApi.DnsrecordDelete(*r.auth, []swagger.DnsRecordRead{{Id: d.Id()}})
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return diags
}

func waitForDnsrecordReady(ctx context.Context, d *schema.ResourceData, r *requester) error {
	return resource.RetryContext(ctx, d.Timeout(schema.TimeoutCreate)-time.Minute, func() *resource.RetryError {
		var err error
		dnsrecordResponse, _, err := r.client.DnsrecordApi.DnsrecordRead(*r.auth, d.Id())
		if err != nil {
			return resource.NonRetryableError(err)
		}

		if !dnsrecordResponse.Ready {
			return resource.RetryableError(fmt.Errorf("not ready yet"))
		}

		return nil
	})
}
