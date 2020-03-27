package gmail

import (
	"net/http"
	"regexp"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"google.golang.org/api/googleapi"
)

func is404Error(err error) bool {
	apiError, ok := err.(*googleapi.Error)
	return ok && apiError.Code == http.StatusNotFound
}

func noTrailingWhitespace() schema.SchemaValidateFunc {
	trailingWhitespace, err := regexp.Compile("[[:space:]]$")
	if err != nil {
		panic(err)
	}
	return validation.StringDoesNotMatch(trailingWhitespace, "name must not end with whitespace")
}
