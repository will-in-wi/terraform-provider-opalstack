package opalstack

import (
	"crypto/sha256"
	"fmt"
	"strings"

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
