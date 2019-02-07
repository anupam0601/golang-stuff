package filescreatormicro

import (
	"context"
	"errors"
	"strings"

	"github.com/go-kit/kit/endpoint"
)

var (
	//ErrRequestTypeNotFound custom error
	ErrRequestTypeNotFound = errors.New("Request type only valid for 'anupam'")
)

//FilesCreateRequest request
type FilesCreateRequest struct {
	RequestType string
	Param       string
}

//FilesCreateResponse response
type FilesCreateResponse struct {
	Message string `json:"message"`
	Err     error  `json:"err,omitempty"`
}

//Endpoints wrapper
type Endpoints struct {
	FilesCreateEndpoint endpoint.Endpoint
}

//MakeFilesCreateEndpoint Creating Files Create Endpoint
//svc is parameter referring to FilesCreatorService
func MakeFilesCreateEndpoint(svc FilesCreatorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(FilesCreateRequest)
		var (
			txt   string
			param string
		)

		param = req.Param

		if strings.EqualFold(req.RequestType, "Word") {
			//Word referring to Word method of our interface FilesCreatorService
			txt = svc.Word(param)
		} else {
			return nil, ErrRequestTypeNotFound
		}
		return FilesCreateResponse{Message: txt}, nil
	}
}
