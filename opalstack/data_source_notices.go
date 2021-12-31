package opalstack

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNotices() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceNoticesRead,
		Schema: map[string]*schema.Schema{
			"notices": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
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
				},
			},
		},
	}
}

func dataSourceNoticesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	r := m.(*requester)

	var diags diag.Diagnostics

	noticeResponse, _, err := r.client.NoticeApi.NoticeList(*r.auth)
	if err != nil {
		return handleSwaggerError(err)
	}

	notices := make([]map[string]interface{}, 0, len(noticeResponse))
	uuids := make([]string, 0, len(noticeResponse))
	for _, elem := range noticeResponse {
		uuids = append(uuids, elem.Id)
		notices = append(notices, map[string]interface{}{
			"id":         elem.Id,
			"type":       elem.Type_,
			"content":    elem.Content,
			"created_at": elem.CreatedAt.Format(time.RFC3339),
		})
	}

	d.SetId(generateIdFromList(uuids))
	if err := d.Set("notices", notices); err != nil {
		return diag.Errorf("Unable to assign notices: %s", err)
	}

	return diags
}
