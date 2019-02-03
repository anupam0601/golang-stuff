package endpoint

import (
	"context"

	service "github.com/anupam0601/golang-stuff/ints-perf-automation/files_creator/pkg/service"
	endpoint "github.com/go-kit/kit/endpoint"
)

// CreateFilesRequest collects the request parameters for the CreateFiles method.
type CreateFilesRequest struct {
	FileType string `json:"file_type"`
}

// CreateFilesResponse collects the response parameters for the CreateFiles method.
type CreateFilesResponse struct {
	S0 string `json:"s0"`
}

// MakeCreateFilesEndpoint returns an endpoint that invokes CreateFiles on the service.
func MakeCreateFilesEndpoint(s service.FilesCreatorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateFilesRequest)
		s0 := s.CreateFiles(ctx, req.FileType)
		return CreateFilesResponse{S0: s0}, nil
	}
}

// CreateFiles implements Service. Primarily useful in a client.
func (e Endpoints) CreateFiles(ctx context.Context, fileType string) (s0 string) {
	request := CreateFilesRequest{FileType: fileType}
	response, err := e.CreateFilesEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(CreateFilesResponse).S0
}
