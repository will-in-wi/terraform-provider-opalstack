package opalstack

import (
	"terraform-provider-opalstack/swagger"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func sitesSchema(idRequired bool) map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Type:     schema.TypeString,
			Computed: !idRequired,
			Required: idRequired,
		},
		"state": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"ready": {
			Type:     schema.TypeBool,
			Computed: true,
		},
		"name": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"server": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"ip4": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"ip6": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"disabled": {
			Type:     schema.TypeBool,
			Computed: true,
		},
		"domains": {
			Type:     schema.TypeSet,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"routes": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"id": {
						Type:     schema.TypeString,
						Computed: true,
					},
					"app": {
						Type:     schema.TypeString,
						Computed: true,
					},
					"uri": {
						Type:     schema.TypeString,
						Computed: true,
					},
				},
			},
		},
		"generate_le": {
			Type:     schema.TypeBool,
			Computed: true,
		},
		"cert": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"redirect": {
			Type:     schema.TypeBool,
			Computed: true,
		},
		"le_http_challenge_tokens": {
			Type:     schema.TypeSet,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
	}
}

func flattenSiteRoutes(routes []swagger.RouteResponse) []map[string]string {
	results := make([]map[string]string, 0)

	for _, route := range routes {
		results = append(results, map[string]string{
			"id":  route.Id,
			"app": route.App,
			"uri": route.Uri,
		})
	}

	return results
}
