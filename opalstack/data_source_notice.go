package opalstack

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNotice() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceNoticeRead,
		Schema: map[string]*schema.Schema{
			"uuid": {
				Type:     schema.TypeString,
				Computed: true,
			},
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

	uuid, ok := d.GetOk("uuid")
	if !ok {
		return diag.Errorf("UUID for Notice is blank")
	}
	strUuid := uuid.(string)

	noticeResponse, _, err := r.client.NoticeApi.NoticeRead(*r.auth, strUuid)
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("uuid", noticeResponse.Id)
	d.Set("type", noticeResponse.Type_)
	d.Set("content", noticeResponse.Content)
	d.Set("created_at", noticeResponse.CreatedAt)
	d.SetId(noticeResponse.Id)

	return diag.Diagnostics{}
}
