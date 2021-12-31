package opalstack

import (
	"fmt"
	"strings"
)

func validateStringInList(arr []string) func(val interface{}, key string) (warns []string, errs []error) {
	return func(val interface{}, key string) (warns []string, errs []error) {
		v := val.(string)

		if !arrayContains(arr, v) {
			errs = append(errs, fmt.Errorf("%q must be one of (%s), got: %s", key, strings.Join(arr, ", "), v))
		}

		return
	}
}
