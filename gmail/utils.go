package gmail

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"google.golang.org/api/googleapi"
)

func is404Error(err error) bool {
	apiError, ok := err.(*googleapi.Error)
	return ok && apiError.Code == http.StatusNotFound
}

// noLeadingTrailingRepeatedSpaces returns a SchemaValidateFunc which tests if
// the provided value does not contain leading, trailing or repeated
// whitespaces.
// It matches space characters as defined by Unicode's White Space property; in
// the Latin-1 space.
func noLeadingTrailingRepeatedSpaces() schema.SchemaValidateFunc {
	const w = "\t\n\v\f\r \u00A0\u0085"
	re := regexp.MustCompile(fmt.Sprintf("^[%s]|[%s][%s]|[%s]$", w, w, w, w))
	return validation.StringDoesNotMatch(re, "name must not contain leading, trailing or repeated whitespaces")
}

// ignoreCase is a SchemaDiffSuppressFunc that compare strings
// case-insensitivity.
func ignoreCase(_, old, new string, _ *schema.ResourceData) bool {
	return strings.EqualFold(old, new)
}

// ignoreLeadingTrailingRepeatedSpaces is a SchemaDiffSuppressFunc that compare
// strings ignoring leading, trailing, and repeated whitespaces.
// It matches space characters as defined by Unicode's White Space property; in
// the Latin-1 space.
func ignoreLeadingTrailingRepeatedSpaces(_, old, new string, _ *schema.ResourceData) bool {
	re := regexp.MustCompile("[\t\n\v\f\r \u00A0\u0085]+")
	o := re.ReplaceAllString(strings.TrimSpace(old), " ")
	n := re.ReplaceAllString(strings.TrimSpace(new), " ")
	return o == n
}
