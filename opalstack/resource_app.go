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

func allowedAppTypes() []string {
	return []string{
		// See swagger/model_app_type_enum.go for list.
		// TODO: Figure out what these mean.
		"STA", // Static Only
		"NPF", // Nginx/PHP-FPM
		"APA", // Apache/PHP-CGI
		"CUS", // Proxied port
		"SLS", // Symlink something?
		"SLP", // Symlink something?
		"SVN", // Subversion
		"DAV", // WebDAV
	}
}

func resourceApp() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceAppCreate,
		ReadContext:   resourceAppRead,
		UpdateContext: resourceAppUpdate,
		DeleteContext: resourceAppDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"osuser": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"type": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateStringInList(allowedAppTypes()),
			},
			"installer_url": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"json": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"server": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"port": {
				Type:     schema.TypeInt,
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

func resourceAppCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	r := m.(*requester)

	appType := swagger.AppTypeEnum(d.Get("type").(string))

	create := swagger.ApplicationCreate{
		Name:         d.Get("name").(string),
		Osuser:       d.Get("osuser").(string),
		Type_:        &appType,
		InstallerUrl: d.Get("installer_url").(string),
		Json:         jsonToStringMap(d.Get("json").(map[string]interface{})),
	}

	appResponse, _, err := r.client.AppApi.AppCreate(*r.auth, []swagger.ApplicationCreate{create})
	if err != nil {
		return handleSwaggerError(err)
	}

	d.SetId(appResponse[0].Id)

	retryErr := waitForAppReady(ctx, d, r)
	if retryErr != nil {
		return diag.Errorf("failed with error while waiting for app to be created: %s", retryErr)
	}

	resourceAppRead(ctx, d, m)

	return diag.Diagnostics{}
}

func resourceAppRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	r := m.(*requester)

	var diags diag.Diagnostics

	appResponse, _, err := r.client.AppApi.AppRead(*r.auth, d.Id())
	if err != nil {
		return handleSwaggerError(err)
	}

	d.Set("name", appResponse.Name)
	d.Set("osuser", appResponse.Osuser)
	d.Set("type", appResponse.Type_)
	d.Set("installer_url", appResponse.InstallerUrl)
	d.Set("json", jsonStructToFlatMap(*appResponse.Json))
	d.Set("server", appResponse.Server)
	d.Set("port", appResponse.Port)

	return diags
}

func resourceAppUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	r := m.(*requester)

	if d.HasChangesExcept("last_updated") {
		appType := swagger.AppTypeEnum(d.Get("type").(string))

		update := swagger.ApplicationUpdate{
			Id:    d.Id(),
			Type_: &appType,
			Json:  jsonToStringMap(d.Get("json").(map[string]interface{})),
		}

		_, _, err := r.client.AppApi.AppUpdate(*r.auth, []swagger.ApplicationUpdate{update})
		if err != nil {
			return handleSwaggerError(err)
		}

		d.Set("last_updated", time.Now().Format(time.RFC850))

		retryErr := waitForAppReady(ctx, d, r)
		if retryErr != nil {
			return diag.Errorf("failed with error while waiting for app to be updated: %s", retryErr)
		}
	}

	return resourceAppRead(ctx, d, m)
}

func resourceAppDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	r := m.(*requester)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	_, err := r.client.AppApi.AppDelete(*r.auth, []swagger.ApplicationRead{{Id: d.Id()}})
	if err != nil {
		return handleSwaggerError(err)
	}

	d.SetId("")

	return diags
}

func waitForAppReady(ctx context.Context, d *schema.ResourceData, r *requester) error {
	return resource.RetryContext(ctx, d.Timeout(schema.TimeoutCreate)-time.Minute, func() *resource.RetryError {
		var err error
		appResponse, _, err := r.client.AppApi.AppRead(*r.auth, d.Id())
		if err != nil {
			return resource.NonRetryableError(err)
		}

		if !appResponse.Ready {
			return resource.RetryableError(fmt.Errorf("not ready yet"))
		}

		return nil
	})
}
