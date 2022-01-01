package opalstack

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

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

func jsonStructToFlatMap(json interface{}) map[string]string {
	st := reflect.TypeOf(json)
	if st.Kind() != reflect.Struct {
		panic(fmt.Sprintf("Type being flat mapped is not struct, it is %s", st.Kind().String()))
	}
	fields := make(map[string]string, 0)
	value := reflect.ValueOf(json)

	for i := 0; i < st.NumField(); i++ {
		field := st.Field(i)
		tag := field.Tag.Get("json")
		if tag != "" {
			parts := strings.Split(tag, ",")
			omitempty := arrayContains(parts[1:], "omitempty")
			fieldName := parts[0]
			fieldValue := valueToString(value.Field(i))
			if fieldName != "" && (!omitempty || fieldValue != "") {
				fields[fieldName] = fieldValue
			}
		}
	}

	return fields
}

func jsonToStringMap(json map[string]interface{}) map[string]string {
	result := make(map[string]string)

	for k, v := range json {
		val := reflect.ValueOf(v)
		result[k] = valueToString(val)
	}

	return result
}

func valueToString(v reflect.Value) string {
	switch v.Type().Kind() {
	case reflect.String:
		return v.String()
	case reflect.Int32, reflect.Int, reflect.Int16, reflect.Int8, reflect.Int64:
		return strconv.Itoa(int(v.Int()))
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())
	default:
		panic(fmt.Sprintf("Unhandled type: %s", v.Type().Name()))
	}
}
