package opalstack

import (
	"context"
	"terraform-provider-opalstack/swagger"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
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
				ValidateFunc: validation.StringInSlice(allowedAppTypes(), false),
			},
			"installer_url": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"json": {
				Type:     schema.TypeList,
				Optional: true,
				MinItems: 1,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"db_name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"db_user": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"db_host": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"db_port": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"app_name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"app_port": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"app_path": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"app_version": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"app_command": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"app_lang_version": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sym_link_path": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"auto_site_url": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"app_exec": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"url_fopen": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"fpm_type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"fpm_max_children": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"fpm_max_requests": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"fpm_start_servers": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"fpm_min_spare_servers": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"fpm_max_spare_servers": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"php_version": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
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
		Json:         parseOutJson(d),
	}

	appResponse, _, err := r.client.AppApi.AppCreate(*r.auth, []swagger.ApplicationCreate{create})
	if err != nil {
		return handleSwaggerError(err)
	}

	d.SetId(appResponse[0].Id)

	retryErr := waitForResourceReady(ctx, d, appChecker(r, d))
	if retryErr != nil {
		return diag.Errorf("failed with error while waiting for app to be created: %s", retryErr)
	}

	resourceAppRead(ctx, d, m)

	return diag.Diagnostics{}
}

func resourceAppRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	r := m.(*requester)

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

	return diag.Diagnostics{}
}

func resourceAppUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	r := m.(*requester)

	if d.HasChangesExcept("last_updated") {
		appType := swagger.AppTypeEnum(d.Get("type").(string))

		update := swagger.ApplicationUpdate{
			Id:    d.Id(),
			Type_: &appType,
			Json:  parseOutJson(d),
		}

		_, _, err := r.client.AppApi.AppUpdate(*r.auth, []swagger.ApplicationUpdate{update})
		if err != nil {
			return handleSwaggerError(err)
		}

		d.Set("last_updated", time.Now().Format(time.RFC850))

		retryErr := waitForResourceReady(ctx, d, appChecker(r, d))
		if retryErr != nil {
			return diag.Errorf("failed with error while waiting for app to be updated: %s", retryErr)
		}
	}

	return resourceAppRead(ctx, d, m)
}

func resourceAppDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	r := m.(*requester)

	_, err := r.client.AppApi.AppDelete(*r.auth, []swagger.ApplicationRead{{Id: d.Id()}})
	if err != nil {
		return handleSwaggerError(err)
	}

	retryErr := waitForResourceDestroyed(ctx, d, appChecker(r, d))
	if retryErr != nil {
		return diag.Errorf("failed with error while waiting for mariauser to be destroyed: %s", retryErr)
	}

	d.SetId("")

	return diag.Diagnostics{}
}

func jsonNames() []string {
	return []string{
		"db_name",
		"db_user",
		"db_host",
		"db_port",
		"app_name",
		"app_port",
		"app_path",
		"app_version",
		"app_command",
		"app_lang_version",
		"sym_link_path",
		"auto_site_url",
		"app_exec",
		"url_fopen",
		"fpm_type",
		"fpm_max_children",
		"fpm_max_requests",
		"fpm_start_servers",
		"fpm_min_spare_servers",
		"fpm_max_spare_servers",
		"php_version",
	}
}

func parseOutJson(d *schema.ResourceData) map[string]interface{} {
	jsonOut := make(map[string]interface{})

	for _, fieldName := range jsonNames() {
		val, ok := d.GetOk("json.0." + fieldName)
		if ok {
			jsonOut[fieldName] = val
		}
	}

	return jsonOut
}

func appChecker(r *requester, d *schema.ResourceData) func() (bool, error) {
	return func() (bool, error) {
		appResponse, _, err := r.client.AppApi.AppRead(*r.auth, d.Id())
		return appResponse.Ready, err
	}
}
