package endpoint

import (
	"context"

	service "github.com/anupam0601/golang-stuff/ints-perf-automation/files_creator/pkg/service"
	endpoint "github.com/go-kit/kit/endpoint"
)

// CreateRequest collects the request parameters for the Create method.
type CreateRequest struct {
	FileDescriptor service.FileDescriptor `json:"file_descriptor"`
}

// CreateResponse collects the response parameters for the Create method.
type CreateResponse struct {
	S0 string `json:"s0"`
}

// MakeCreateEndpoint returns an endpoint that invokes Create on the service.
func MakeCreateEndpoint(s service.FilesCreatorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateRequest)
		s0 := s.Create(ctx, req.FileDescriptor)
		return CreateResponse{S0: s0}, nil
	}
}

// Create implements Service. Primarily useful in a client.
func (e Endpoints) Create(ctx context.Context, fileDescriptor service.FileDescriptor) (s0 string) {
	request := CreateRequest{FileDescriptor: fileDescriptor}
	response, err := e.CreateEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(CreateResponse).S0
}
