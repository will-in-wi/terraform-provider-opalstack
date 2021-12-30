package opalstack

import (
	"context"
	"terraform-provider-opalstack/swagger"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"base_path": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OPALSTACK_BASE_PATH", "https://my.opalstack.com"),
			},
			"token": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("OPALSTACK_TOKEN", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			// "hashicups_order": resourceOrder(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"opalstack_cert":        dataSourceCert(),
			"opalstack_certs":       dataSourceCerts(),
			"opalstack_cert_shared": dataSourceCertShared(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

type requester struct {
	client *swagger.APIClient
	auth   *context.Context
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	token := d.Get("token").(string)

	auth := context.WithValue(context.Background(), swagger.ContextAPIKey, swagger.APIKey{
		Prefix: "token",
		Key:    token,
	})

	var basePath *string

	hVal, ok := d.GetOk("base_path")
	if ok {
		tempHost := hVal.(string)
		basePath = &tempHost
	}

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	cfg := swagger.NewConfiguration()
	cfg.UserAgent = "terraform-provider-opalstack/0.0.1"
	if basePath == nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Missing base_path",
			Detail:   "The base_path parameter has been set to nothing",
		})
		return nil, diags
	}
	cfg.BasePath = *basePath

	c := swagger.NewAPIClient(cfg)

	return &requester{c, &auth}, diags
}
