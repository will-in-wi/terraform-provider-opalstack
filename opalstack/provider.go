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
			"opalstack_notice":    resourceNotice(),
			"opalstack_domain":    resourceDomain(),
			"opalstack_dnsrecord": resourceDnsrecord(),
			"opalstack_osuser":    resourceOsuser(),
			"opalstack_osvar":     resourceOsvar(),
			"opalstack_app":       resourceApp(),
			"opalstack_psqluser":  resourcePsqluser(),
			"opalstack_psqldb":    resourcePsqldb(),
			"opalstack_mariauser": resourceMariauser(),
			"opalstack_mariadb":   resourceMariadb(),
			"opalstack_site":      resourceSite(),
			"opalstack_cert":      resourceCert(),
			"opalstack_mailuser":  resourceMailuser(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"opalstack_notices":     dataSourceNotices(),
			"opalstack_notice":      dataSourceNotice(),
			"opalstack_servers":     dataSourceServers(),
			"opalstack_server":      dataSourceServer(),
			"opalstack_ips":         dataSourceIps(),
			"opalstack_ip":          dataSourceIp(),
			"opalstack_domains":     dataSourceDomains(),
			"opalstack_domain":      dataSourceDomain(),
			"opalstack_dnsrecords":  dataSourceDnsrecords(),
			"opalstack_dnsrecord":   dataSourceDnsrecord(),
			"opalstack_account":     dataSourceAccount(),
			"opalstack_osusers":     dataSourceOsusers(),
			"opalstack_osuser":      dataSourceOsuser(),
			"opalstack_osvars":      dataSourceOsvars(),
			"opalstack_osvar":       dataSourceOsvar(),
			"opalstack_psqlusers":   dataSourcePsqlusers(),
			"opalstack_psqluser":    dataSourcePsqluser(),
			"opalstack_psqldbs":     dataSourcePsqldbs(),
			"opalstack_psqldb":      dataSourcePsqldb(),
			"opalstack_mariausers":  dataSourceMariausers(),
			"opalstack_mariauser":   dataSourceMariauser(),
			"opalstack_mariadbs":    dataSourceMariadbs(),
			"opalstack_mariadb":     dataSourceMariadb(),
			"opalstack_apps":        dataSourceApps(),
			"opalstack_app":         dataSourceApp(),
			"opalstack_mailusers":   dataSourceMailusers(),
			"opalstack_mailuser":    dataSourceMailuser(),
			"opalstack_sites":       dataSourceSites(),
			"opalstack_site":        dataSourceSite(),
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
