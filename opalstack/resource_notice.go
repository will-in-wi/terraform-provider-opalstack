package opalstack

import (
	"context"
	"terraform-provider-opalstack/swagger"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func allowedTypes() []string {
	return []string{
		// See swagger/model_notice_type_enum.go for list.
		"M", // Message
		"P", // Password Change
		"D", // Default Password
		"R", // Resource Overage
	}
}

func resourceNotice() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceNoticeCreate,
		ReadContext:   resourceNoticeRead,
		UpdateContext: resourceNoticeUpdate,
		DeleteContext: resourceNoticeDelete,
		Schema: map[string]*schema.Schema{
			"type": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validateStringInList(allowedTypes()),
			},
			"content": {
				Type:     schema.TypeString,
				Required: true,
			},
			"created_at": {
				Type:     schema.TypeString,
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

func resourceNoticeCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	r := m.(*requester)

	noticeType := swagger.NoticeTypeEnum(d.Get("type").(string))

	body := []swagger.NoticeCreate{
		{
			Type_:   &noticeType,
			Content: d.Get("content").(string),
		},
	}

	certResponse, _, err := r.client.NoticeApi.NoticeCreate(*r.auth, body)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(certResponse[0].Id)
	resourceNoticeRead(ctx, d, m)

	return diag.Diagnostics{}
}

func resourceNoticeRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	r := m.(*requester)

	var diags diag.Diagnostics

	noticeResponse, _, err := r.client.NoticeApi.NoticeRead(*r.auth, d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("type", noticeResponse.Type_)
	d.Set("content", noticeResponse.Content)
	d.Set("created_at", noticeResponse.CreatedAt.Format(time.RFC3339))

	return diags
}

func resourceNoticeUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	r := m.(*requester)

	if d.HasChangesExcept("last_updated") {
		noticeType := swagger.NoticeTypeEnum(d.Get("type").(string))
		body := []swagger.NoticeUpdate{
			{
				Id:      d.Id(),
				Type_:   &noticeType,
				Content: d.Get("content").(string),
			},
		}

		_, _, err := r.client.NoticeApi.NoticeUpdate(*r.auth, body)
		if err != nil {
			return diag.FromErr(err)
		}

		d.Set("last_updated", time.Now().Format(time.RFC850))
	}

	return resourceNoticeRead(ctx, d, m)
}

func resourceNoticeDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	r := m.(*requester)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	_, err := r.client.NoticeApi.NoticeDelete(*r.auth, []swagger.NoticeRead{{Id: d.Id()}})
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return diags
}
