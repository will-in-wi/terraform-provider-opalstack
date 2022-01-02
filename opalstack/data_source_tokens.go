package opalstack

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceTokens() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceTokensRead,
		Schema: map[string]*schema.Schema{
			"tokens": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Resource{Schema: tokensSchema(false)},
			},
		},
	}
}

func dataSourceTokensRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	r := m.(*requester)

	var diags diag.Diagnostics

	tokensResponse, _, err := r.client.TokenApi.TokenList(*r.auth)
	if err != nil {
		return handleSwaggerError(err)
	}

	uuids := make([]string, 0, len(tokensResponse))
	tokens := make([]map[string]interface{}, 0, len(tokensResponse))
	for _, elem := range tokensResponse {
		uuids = append(uuids, elem.Key)
		tokens = append(tokens, map[string]interface{}{
			"id":   elem.Key,
			"name": elem.Name,
			"key":  elem.Key,
		})
	}
	if err := d.Set("tokens", tokens); err != nil {
		return diag.Errorf("Unable to assign tokens: %s", err)
	}

	d.SetId(generateIdFromList(uuids))

	return diags
}
