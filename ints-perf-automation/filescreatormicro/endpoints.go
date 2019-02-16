package filescreatormicro

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

//FilesCreateRequest request
type FilesCreateRequest struct {
	FileDescriptor FileDescriptor `json:"file_descriptor"`
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
		msg, err := svc.Create(req.FileDescriptor)
		return FilesCreateResponse{Message:msg, Err:err}, nil
		//var (
		//	txt   string
		//	param string
		//)
		//
		//param = req.Param
		//
		//if strings.EqualFold(req.RequestType, "Word") {
		//	//Word referring to Word method of our interface FilesCreatorService
		//	txt = svc.Word(param)
		//} else {
		//	return nil, ErrRequestTypeNotFound
		//}
		//return FilesCreateResponse{Message: txt}, nil

	}
}
