package opalstack

import (
	"reflect"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestCompareTrimmed(t *testing.T) {
	k := ""
	d := &schema.ResourceData{}
	if !compareTrimmed(k, "sjf ", "sjf", d) {
		t.Errorf("Something is broken...")
	}

	if !compareTrimmed(k, "\nsjf\n", "sjf", d) {
		t.Errorf("Something is broken...")
	}
}

func TestGenerateIdFromList(t *testing.T) {
	res := generateIdFromList([]string{"1", "2", "3"})
	if res != "a665a45920422f9d417e4867efdc4fb8a04a1f3fff1fa07e998e86f7f7a27ae3" {
		t.Errorf("This isn't good")
	}
}

func TestStringSetToStringArrayRoundTrip(t *testing.T) {
	expected := []string{"hello", "world"}
	input := stringArrayToStringSet(expected)
	res := stringSetToStringArray(input)
	if !reflect.DeepEqual(res, expected) {
		t.Errorf("Danger, Will Robinson!")
	}
}
