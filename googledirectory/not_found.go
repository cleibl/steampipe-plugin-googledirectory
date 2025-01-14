package googledirectory

import (
	"github.com/turbot/go-kit/helpers"
	"github.com/turbot/go-kit/types"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"google.golang.org/api/googleapi"
)

// function which returns an IsNotFoundErrorPredicate for Google Directory API calls
func isNotFoundError(notFoundErrors []string) plugin.ErrorPredicate {
	return func(err error) bool {
		if gerr, ok := err.(*googleapi.Error); ok {
			return helpers.StringSliceContains(notFoundErrors, types.ToString(gerr.Code))
		}
		return false
	}
}
