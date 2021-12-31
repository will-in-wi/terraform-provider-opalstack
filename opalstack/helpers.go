package opalstack

import (
	"crypto/sha256"
	"fmt"
	"strings"
	"terraform-provider-opalstack/swagger"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Compare while ignoring whitespace.
func compareTrimmed(k, old, new string, d *schema.ResourceData) bool {
	return strings.TrimSpace(old) == strings.TrimSpace(new)
}

func generateIdFromList(strs []string) string {
	joinedString := strings.Join(strs, "")
	hash := sha256.Sum256([]byte(joinedString))
	return fmt.Sprintf("%x", hash)
}

func arrayContains(haystack []string, needle string) bool {
	for _, s := range haystack {
		if s == needle {
			return true
		}
	}

	return false
}

func stringSetToStringArray(set *schema.Set) []string {
	result := make([]string, 0)

	for _, v := range set.List() {
		val, ok := v.(string)
		if ok && val != "" {
			result = append(result, val)
		}
	}

	return result
}

func stringArrayToStringSet(strs []string) *schema.Set {
	vs := make([]interface{}, 0, len(strs))
	for _, v := range strs {
		vs = append(vs, v)
	}
	return schema.NewSet(schema.HashString, vs)
}

func handleSwaggerError(err error) diag.Diagnostics {
	swaggerError, ok := err.(swagger.GenericSwaggerError)
	if ok {
		return diag.Diagnostics{
			diag.Diagnostic{
				Severity: diag.Error,
				Summary:  fmt.Sprintf("Server returned an invalid response: %s", swaggerError.Error()),
				Detail:   string(swaggerError.Body()),
			},
		}
	} else {
		return diag.FromErr(err)
	}
}
