package opalstack

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNotice() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceNoticeRead,
		Schema: map[string]*schema.Schema{
			"type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"content": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"created_at": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceNoticeRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	r := m.(*requester)

	noticeResponse, _, err := r.client.NoticeApi.NoticeRead(*r.auth, d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("type", noticeResponse.Type_)
	d.Set("content", noticeResponse.Content)
	d.Set("created_at", noticeResponse.CreatedAt.Format(time.RFC3339))
	d.SetId(noticeResponse.Id)

	return diag.Diagnostics{}
}
