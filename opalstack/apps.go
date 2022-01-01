package opalstack

import (
	"reflect"
	"strings"
	"terraform-provider-opalstack/swagger"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func appsSchema(idRequired bool) map[string]*schema.Schema {
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
		"osuser": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"type": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"port": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"installer_url": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"json": {
			Type:     schema.TypeMap,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
	}
}

func jsonStructToFlatMap(json swagger.ApplicationResponseJson) map[string]interface{} {
	st := reflect.TypeOf(json)
	fields := make(map[string]interface{}, 0)
	value := reflect.ValueOf(json)

	for i := 0; i < st.NumField(); i++ {
		field := st.Field(i)
		tag := field.Tag.Get("json")
		if tag != "" {
			parts := strings.Split(tag, ",")
			fieldName := parts[0]
			field := value.Field(i)
			if field.Type().Kind() != reflect.Ptr || !field.IsNil() {
				fields[fieldName] = field.Interface()
			}
		}
	}

	return fields
}
