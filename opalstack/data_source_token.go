package opalstack

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceToken() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceTokenRead,
		Schema:      tokensSchema(true),
	}
}

func dataSourceTokenRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	r := m.(*requester)

	var diags diag.Diagnostics

	tokenResponse, _, err := r.client.TokenApi.TokenRead(*r.auth, d.Get("id").(string))
	if err != nil {
		return handleSwaggerError(err)
	}

	d.Set("name", tokenResponse.Name)
	d.Set("key", tokenResponse.Key)
	d.SetId(tokenResponse.Key)

	return diags
}
