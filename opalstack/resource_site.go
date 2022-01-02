package opalstack

import (
	"context"
	"terraform-provider-opalstack/swagger"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSite() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSiteCreate,
		ReadContext:   resourceSiteRead,
		UpdateContext: resourceSiteUpdate,
		DeleteContext: resourceSiteDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"ip4": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ip6": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"domains": {
				Type:     schema.TypeSet,
				Required: true,
				MinItems: 1,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"routes": {
				Type:     schema.TypeList,
				Required: true,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uri": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"app": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"cert": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"redirect": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"generate_le": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"disabled": {
				Type:     schema.TypeBool,
				Optional: true,
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

func resourceSiteCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	r := m.(*requester)

	create := swagger.SiteCreate{
		Name:       d.Get("name").(string),
		Ip4:        d.Get("ip4").(string),
		Ip6:        d.Get("ip6").(string),
		Domains:    stringSetToStringArray(d.Get("domains").(*schema.Set)),
		Routes:     expandSiteRoutes(d.Get("routes").([]interface{})),
		Cert:       d.Get("cert").(string),
		Redirect:   d.Get("redirect").(bool),
		GenerateLe: d.Get("generate_le").(bool),
		Disabled:   d.Get("disabled").(bool),
	}

	siteResponse, _, err := r.client.SiteApi.SiteCreate(*r.auth, []swagger.SiteCreate{create})
	if err != nil {
		return handleSwaggerError(err)
	}

	d.SetId(siteResponse[0].Id)

	retryErr := waitForResourceReady(ctx, d, siteChecker(r, d))
	if retryErr != nil {
		return diag.Errorf("failed with error while waiting for site to be created: %s", retryErr)
	}

	resourceSiteRead(ctx, d, m)

	return diag.Diagnostics{}
}

func resourceSiteRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	r := m.(*requester)

	siteResponse, _, err := r.client.SiteApi.SiteRead(*r.auth, d.Id())
	if err != nil {
		return handleSwaggerError(err)
	}

	d.Set("name", siteResponse.Name)
	d.Set("ip4", siteResponse.Ip4)
	d.Set("ip6", siteResponse.Ip6)
	d.Set("domains", siteResponse.Domains)
	d.Set("routes", siteResponse.Routes)
	d.Set("cert", siteResponse.Cert)
	d.Set("redirect", siteResponse.Redirect)
	d.Set("generate_le", siteResponse.GenerateLe)
	d.Set("disabled", siteResponse.Disabled)

	return diag.Diagnostics{}
}

func resourceSiteUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	r := m.(*requester)

	if d.HasChangesExcept("last_updated") {
		update := swagger.SiteUpdate{
			Id:         d.Id(),
			Name:       d.Get("name").(string),
			Ip4:        d.Get("ip4").(string),
			Ip6:        d.Get("ip6").(string),
			Domains:    stringSetToStringArray(d.Get("domains").(*schema.Set)),
			Routes:     expandSiteRoutes(d.Get("routes").([]interface{})),
			Cert:       d.Get("cert").(string),
			Redirect:   d.Get("redirect").(bool),
			GenerateLe: d.Get("generate_le").(bool),
			Disabled:   d.Get("disabled").(bool),
		}

		_, _, err := r.client.SiteApi.SiteUpdate(*r.auth, []swagger.SiteUpdate{update})
		if err != nil {
			return handleSwaggerError(err)
		}

		d.Set("last_updated", time.Now().Format(time.RFC850))

		retryErr := waitForResourceReady(ctx, d, siteChecker(r, d))
		if retryErr != nil {
			return diag.Errorf("failed with error while waiting for site to be updated: %s", retryErr)
		}
	}

	return resourceSiteRead(ctx, d, m)
}

func resourceSiteDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	r := m.(*requester)

	_, err := r.client.SiteApi.SiteDelete(*r.auth, []swagger.SiteRead{{Id: d.Id()}})
	if err != nil {
		return handleSwaggerError(err)
	}

	retryErr := waitForResourceDestroyed(ctx, d, siteChecker(r, d))
	if retryErr != nil {
		return diag.Errorf("failed with error while waiting for site to be destroyed: %s", retryErr)
	}

	d.SetId("")

	return diag.Diagnostics{}
}

func expandSiteRoutes(routes []interface{}) []swagger.Route {
	results := make([]swagger.Route, 0)

	for _, route := range routes {
		routeStruct := route.(map[string]interface{})
		results = append(results, swagger.Route{
			Uri: routeStruct["uri"].(string),
			App: routeStruct["app"].(string),
		})
	}

	return results
}

func siteChecker(r *requester, d *schema.ResourceData) func() (bool, error) {
	return func() (bool, error) {
		siteResponse, _, err := r.client.SiteApi.SiteRead(*r.auth, d.Id())
		return siteResponse.Ready, err
	}
}
