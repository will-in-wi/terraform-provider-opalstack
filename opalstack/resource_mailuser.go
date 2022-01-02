package opalstack

import (
	"context"
	"terraform-provider-opalstack/swagger"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
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

	retryErr := waitForResourceReady(ctx, d, mailuserChecker(r, d))
	if retryErr != nil {
		return diag.Errorf("failed with error while waiting for mailuser to be created: %s", retryErr)
	}

	resourceMailuserRead(ctx, d, m)

	return diag.Diagnostics{}
}

func resourceMailuserRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	r := m.(*requester)

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

	return diag.Diagnostics{}
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

		retryErr := waitForResourceReady(ctx, d, mailuserChecker(r, d))
		if retryErr != nil {
			return diag.Errorf("failed with error while waiting for mailuser to be updated: %s", retryErr)
		}
	}

	return resourceMailuserRead(ctx, d, m)
}

func resourceMailuserDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	r := m.(*requester)

	_, err := r.client.MailuserApi.MailuserDelete(*r.auth, []swagger.MailUserDelete{{Id: d.Id()}})
	if err != nil {
		return handleSwaggerError(err)
	}

	retryErr := waitForResourceDestroyed(ctx, d, mailuserChecker(r, d))
	if retryErr != nil {
		return diag.Errorf("failed with error while waiting for mailuser to be destroyed: %s", retryErr)
	}

	d.SetId("")

	return diag.Diagnostics{}
}

func mailuserChecker(r *requester, d *schema.ResourceData) func() (bool, error) {
	return func() (bool, error) {
		mailuserResponse, _, err := r.client.MailuserApi.MailuserRead(*r.auth, d.Id())
		return mailuserResponse.Ready, err
	}
}
