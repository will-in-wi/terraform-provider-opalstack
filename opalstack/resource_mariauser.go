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

func resourceMariauser() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceMariauserCreate,
		ReadContext:   resourceMariauserRead,
		UpdateContext: resourceMariauserUpdate,
		DeleteContext: resourceMariauserDelete,
		Schema: map[string]*schema.Schema{
			"server": {
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
			"external": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceMariauserCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	r := m.(*requester)

	create := swagger.MariaUserCreate{
		Server:   d.Get("server").(string),
		Name:     d.Get("name").(string),
		Password: d.Get("password").(string),
		External: d.Get("external").(bool),
	}

	mariauserResponse, _, err := r.client.MariauserApi.MariauserCreate(*r.auth, []swagger.MariaUserCreate{create})
	if err != nil {
		return handleSwaggerError(err)
	}

	d.SetId(mariauserResponse[0].Id)

	retryErr := waitForMariauserReady(ctx, d, r)
	if retryErr != nil {
		return diag.Errorf("failed with error while waiting for mariauser to be created: %s", retryErr)
	}

	resourceMariauserRead(ctx, d, m)

	return diag.Diagnostics{}
}

func resourceMariauserRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	r := m.(*requester)

	var diags diag.Diagnostics

	mariauserResponse, _, err := r.client.MariauserApi.MariauserRead(*r.auth, d.Id())
	if err != nil {
		return handleSwaggerError(err)
	}

	d.Set("server", mariauserResponse.Server)
	d.Set("name", mariauserResponse.Name)
	d.Set("external", mariauserResponse.External)

	return diags
}

func resourceMariauserUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	r := m.(*requester)

	if d.HasChangesExcept("last_updated") {
		update := swagger.MariaUserUpdate{
			Id:       d.Id(),
			Password: d.Get("password").(string),
			External: d.Get("external").(bool),
		}

		_, _, err := r.client.MariauserApi.MariauserUpdate(*r.auth, []swagger.MariaUserUpdate{update})
		if err != nil {
			return handleSwaggerError(err)
		}

		d.Set("last_updated", time.Now().Format(time.RFC850))

		retryErr := waitForMariauserReady(ctx, d, r)
		if retryErr != nil {
			return diag.Errorf("failed with error while waiting for mariauser to be updated: %s", retryErr)
		}
	}

	return resourceMariauserRead(ctx, d, m)
}

func resourceMariauserDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	r := m.(*requester)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	_, err := r.client.MariauserApi.MariauserDelete(*r.auth, []swagger.MariaUserRead{{Id: d.Id()}})
	if err != nil {
		return handleSwaggerError(err)
	}

	d.SetId("")

	return diags
}

func waitForMariauserReady(ctx context.Context, d *schema.ResourceData, r *requester) error {
	return resource.RetryContext(ctx, d.Timeout(schema.TimeoutCreate)-time.Minute, func() *resource.RetryError {
		var err error
		mariauserResponse, _, err := r.client.MariauserApi.MariauserRead(*r.auth, d.Id())
		if err != nil {
			return resource.NonRetryableError(err)
		}

		if !mariauserResponse.Ready {
			return resource.RetryableError(fmt.Errorf("not ready yet"))
		}

		return nil
	})
}
