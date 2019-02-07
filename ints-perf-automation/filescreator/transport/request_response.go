package transport

import (
	"github.com/anupam0601/golang-stuff/ints-perf-automation/filescreator"
)

// CreateRequest holds the request parameters for the Create method.
type CreateRequest struct {
	FileDescriptor filescreator.FileDescriptor
}

// CreateResponse holds the response values for the Create method.
type CreateResponse struct {
	Name string `json:"name"`
	Err  error  `json:"error,omitempty"`
}
