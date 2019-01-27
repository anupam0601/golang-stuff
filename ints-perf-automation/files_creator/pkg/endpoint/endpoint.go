package endpoint

import (
	"context"

	service "github.com/anupam0601/golang-stuff/ints-perf-automation/files_creator/pkg/service"
	endpoint "github.com/go-kit/kit/endpoint"
)

// CreateFilesRequest collects the request parameters for the CreateFiles method.
type CreateFilesRequest struct {
	NumberOfFiles int `json:"number_of_files"`
}

// CreateFilesResponse collects the response parameters for the CreateFiles method.
type CreateFilesResponse struct {
	E0 error `json:"e0"`
}

// MakeCreateFilesEndpoint returns an endpoint that invokes CreateFiles on the service.
func MakeCreateFilesEndpoint(s service.FilesCreatorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateFilesRequest)
		e0 := s.CreateFiles(ctx, req.NumberOfFiles)
		return CreateFilesResponse{E0: e0}, nil
	}
}

// Failed implements Failer.
func (r CreateFilesResponse) Failed() error {
	return r.E0
}

// StoreFileMetadataRequest collects the request parameters for the StoreFileMetadata method.
type StoreFileMetadataRequest struct {
	FileName  string `json:"file_name"`
	TimeStamp string `json:"time_stamp"`
}

// StoreFileMetadataResponse collects the response parameters for the StoreFileMetadata method.
type StoreFileMetadataResponse struct {
	E0 error `json:"e0"`
}

// MakeStoreFileMetadataEndpoint returns an endpoint that invokes StoreFileMetadata on the service.
func MakeStoreFileMetadataEndpoint(s service.FilesCreatorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(StoreFileMetadataRequest)
		e0 := s.StoreFileMetadata(ctx, req.FileName, req.TimeStamp)
		return StoreFileMetadataResponse{E0: e0}, nil
	}
}

// Failed implements Failer.
func (r StoreFileMetadataResponse) Failed() error {
	return r.E0
}

// Failer is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// CreateFiles implements Service. Primarily useful in a client.
func (e Endpoints) CreateFiles(ctx context.Context, numberOfFiles int) (e0 error) {
	request := CreateFilesRequest{NumberOfFiles: numberOfFiles}
	response, err := e.CreateFilesEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(CreateFilesResponse).E0
}

// StoreFileMetadata implements Service. Primarily useful in a client.
func (e Endpoints) StoreFileMetadata(ctx context.Context, fileName string, timeStamp string) (e0 error) {
	request := StoreFileMetadataRequest{
		FileName:  fileName,
		TimeStamp: timeStamp,
	}
	response, err := e.StoreFileMetadataEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(StoreFileMetadataResponse).E0
}
