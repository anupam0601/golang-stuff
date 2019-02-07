package transport

import (
	"context"

	"github.com/go-kit/kit/endpoint"

	"github.com/anupam0601/golang-stuff/ints-perf-automation/filescreator"
)

// Endpoints holds all Go kit endpoints for the Order service.
type Endpoints struct {
	Create endpoint.Endpoint
}

// MakeEndpoints initializes all Go kit endpoints for the Order service.
func MakeEndpoints(s filescreator.FilesCreator) Endpoints {
	return Endpoints{
		Create: makeCreateEndpoint(s),
	}
}

func makeCreateEndpoint(s filescreator.FilesCreator) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateRequest)
		name, err := s.Create(ctx, req.FileDescriptor)
		return CreateResponse{Name: name, Err: err}, nil
	}
}
