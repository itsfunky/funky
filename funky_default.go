// +build !aws,!local

// Package funky lets you develop and build lambdas/functions for multiple cloud
// services, while allowing you to test and iterate quickly when developing
// locally.
package funky

import (
	"net/http"
)

// Handle exposes your handler to your cloud function service.
//
// Handle is implemented slightly differently for each cloud provider, but
// attempts to unify each implementation around an HTTP-based handler.
func Handle(handler http.Handler) {
	// This should never be called unless there was a bad build configuration.
	panic("we should never get here")
}
