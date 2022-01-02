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

func resourceMailuser() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceMailuserCreate,
		ReadContext:   resourceMailuserRead,
		UpdateContext: resourceMailuserUpdate,
		DeleteContext: resourceMailuserDelete,
		Schema: map[string]*schema.Schema{
			"imap_server": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"password": {
				Type:      schema.TypeString,
				Optional:  true,
				Sensitive: true,
			},
			"procmailrc": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"autoresponder_enable": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"autoresponder_subject": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"autoresponder_message": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceMailuserCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	r := m.(*requester)

	create := swagger.MailUserCreate{
		ImapServer:           d.Get("imap_server").(string),
		Name:                 d.Get("name").(string),
		Password:             d.Get("password").(string),
		Procmailrc:           d.Get("procmailrc").(string),
		AutoresponderEnable:  d.Get("autoresponder_enable").(bool),
		AutoresponderSubject: d.Get("autoresponder_subject").(string),
		AutoresponderMessage: d.Get("autoresponder_message").(string),
	}

	mailuserResponse, _, err := r.client.MailuserApi.MailuserCreate(*r.auth, []swagger.MailUserCreate{create})
	if err != nil {
		return handleSwaggerError(err)
	}

	d.SetId(mailuserResponse[0].Id)

	retryErr := waitForMailuserReady(ctx, d, r)
	if retryErr != nil {
		return diag.Errorf("failed with error while waiting for mailuser to be created: %s", retryErr)
	}

	resourceMailuserRead(ctx, d, m)

	return diag.Diagnostics{}
}

func resourceMailuserRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	r := m.(*requester)

	var diags diag.Diagnostics

	mailuserResponse, _, err := r.client.MailuserApi.MailuserRead(*r.auth, d.Id())
	if err != nil {
		return handleSwaggerError(err)
	}

	d.Set("imap_server", mailuserResponse.ImapServer)
	d.Set("name", mailuserResponse.Name)
	d.Set("procmailrc", mailuserResponse.Procmailrc)
	d.Set("autoresponder_enable", mailuserResponse.AutoresponderEnable)
	d.Set("autoresponder_subject", mailuserResponse.AutoresponderSubject)
	d.Set("autoresponder_message", mailuserResponse.AutoresponderMessage)

	return diags
}

func resourceMailuserUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	r := m.(*requester)

	if d.HasChangesExcept("last_updated") {
		update := swagger.MailUserUpdate{
			Id:                   d.Id(),
			Password:             d.Get("password").(string),
			Procmailrc:           d.Get("procmailrc").(string),
			AutoresponderEnable:  d.Get("autoresponder_enable").(bool),
			AutoresponderSubject: d.Get("autoresponder_subject").(string),
			AutoresponderMessage: d.Get("autoresponder_message").(string),
		}

		_, _, err := r.client.MailuserApi.MailuserUpdate(*r.auth, []swagger.MailUserUpdate{update})
		if err != nil {
			return handleSwaggerError(err)
		}

		d.Set("last_updated", time.Now().Format(time.RFC850))

		retryErr := waitForMailuserReady(ctx, d, r)
		if retryErr != nil {
			return diag.Errorf("failed with error while waiting for mailuser to be updated: %s", retryErr)
		}
	}

	return resourceMailuserRead(ctx, d, m)
}

func resourceMailuserDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	r := m.(*requester)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	_, err := r.client.MailuserApi.MailuserDelete(*r.auth, []swagger.MailUserDelete{{Id: d.Id()}})
	if err != nil {
		return handleSwaggerError(err)
	}

	d.SetId("")

	return diags
}

func waitForMailuserReady(ctx context.Context, d *schema.ResourceData, r *requester) error {
	return resource.RetryContext(ctx, d.Timeout(schema.TimeoutCreate)-time.Minute, func() *resource.RetryError {
		var err error
		mailuserResponse, _, err := r.client.MailuserApi.MailuserRead(*r.auth, d.Id())
		if err != nil {
			return resource.NonRetryableError(err)
		}

		if !mailuserResponse.Ready {
			return resource.RetryableError(fmt.Errorf("not ready yet"))
		}

		return nil
	})
}
