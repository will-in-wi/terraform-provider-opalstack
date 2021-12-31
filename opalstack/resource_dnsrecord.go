package opalstack

import (
	"context"
	"terraform-provider-opalstack/swagger"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
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
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ready": {
				Type:     schema.TypeBool,
				Computed: true,
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
		Domain:  d.Get("domain").(string),
		Type_:   &dnsrecordType,
		Content: d.Get("content").(string),
	}

	priority, ok := d.GetOk("priority")
	if ok {
		create.Priority = int32(priority.(int))
	}

	ttl, ok := d.GetOk("ttl")
	if ok {
		create.Ttl = int32(ttl.(int))
	}

	dnsrecordResponse, _, err := r.client.DnsrecordApi.DnsrecordCreate(*r.auth, []swagger.DnsRecordCreate{create})
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(dnsrecordResponse[0].Id)
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
	d.Set("state", dnsrecordResponse.State)
	d.Set("ready", dnsrecordResponse.Ready)

	return diags
}

func resourceDnsrecordUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	r := m.(*requester)

	if d.HasChangesExcept("last_updated") {
		dnsrecordType := swagger.DnsRecordTypeEnum(d.Get("type").(string))
		update := swagger.DnsRecordUpdate{
			Id:      d.Id(),
			Domain:  d.Get("domain").(string),
			Type_:   &dnsrecordType,
			Content: d.Get("content").(string),
		}

		priority, ok := d.GetOk("priority")
		if ok {
			update.Priority = int32(priority.(int))
		}

		ttl, ok := d.GetOk("ttl")
		if ok {
			update.Ttl = int32(ttl.(int))
		}

		_, _, err := r.client.DnsrecordApi.DnsrecordUpdate(*r.auth, []swagger.DnsRecordUpdate{update})
		if err != nil {
			return diag.FromErr(err)
		}

		d.Set("last_updated", time.Now().Format(time.RFC850))
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
