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

func generateIdFromList(uuids []string) string {
	joinedString := strings.Join(uuids, "")
	hash := sha256.Sum256([]byte(joinedString))
	return fmt.Sprintf("%x", hash)
}
