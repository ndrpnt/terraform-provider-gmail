package gmail

import (
	"net/http"

	"google.golang.org/api/googleapi"
)

func is404Error(err error) bool {
	apiError, ok := err.(*googleapi.Error)
	return ok && apiError.Code == http.StatusNotFound
}
