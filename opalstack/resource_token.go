package opalstack

import (
	"context"
	"terraform-provider-opalstack/swagger"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceToken() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceTokenCreate,
		ReadContext:   resourceTokenRead,
		UpdateContext: resourceTokenUpdate,
		DeleteContext: resourceTokenDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"key": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceTokenCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	r := m.(*requester)

	create := swagger.TokenCreate{
		Name: d.Get("name").(string),
	}

	tokenResponse, _, err := r.client.TokenApi.TokenCreate(*r.auth, []swagger.TokenCreate{create})
	if err != nil {
		return handleSwaggerError(err)
	}

	d.SetId(tokenResponse[0].Key)

	resourceTokenRead(ctx, d, m)

	return diag.Diagnostics{}
}

func resourceTokenRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	r := m.(*requester)

	tokenResponse, _, err := r.client.TokenApi.TokenRead(*r.auth, d.Id())
	if err != nil {
		return handleSwaggerError(err)
	}

	d.Set("name", tokenResponse.Name)
	d.Set("key", tokenResponse.Key)

	return diag.Diagnostics{}
}

func resourceTokenUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	r := m.(*requester)

	if d.HasChangesExcept("last_updated") {
		update := swagger.TokenUpdate{
			Key:  d.Id(),
			Name: d.Get("name").(string),
		}

		_, _, err := r.client.TokenApi.TokenUpdate(*r.auth, []swagger.TokenUpdate{update})
		if err != nil {
			return handleSwaggerError(err)
		}

		d.Set("last_updated", time.Now().Format(time.RFC850))
	}

	return resourceTokenRead(ctx, d, m)
}

func resourceTokenDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	r := m.(*requester)

	_, err := r.client.TokenApi.TokenDelete(*r.auth, []swagger.TokenRead{{Key: d.Id()}})
	if err != nil {
		return handleSwaggerError(err)
	}

	d.SetId("")

	return diag.Diagnostics{}
}
