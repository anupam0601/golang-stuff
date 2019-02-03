package service

import (
	"context"
	"log"
)

// FilesCreatorService describes the service.
type FilesCreatorService interface {
	// Add your methods here
	// e.x: Foo(ctx context.Context,s string)(rs string, err error)
	CreateFiles(ctx context.Context, fileType string) string
}

type basicFilesCreatorService struct{}

func (b *basicFilesCreatorService) CreateFiles(ctx context.Context, fileType string) (s0 string) {
	// TODO implement the business logic of CreateFiles
	if fileType == "txt" {
		log.Printf("Creating file of type: %s", fileType)
	}
	return fileType
}

// NewBasicFilesCreatorService returns a naive, stateless implementation of FilesCreatorService.
func NewBasicFilesCreatorService() FilesCreatorService {
	return &basicFilesCreatorService{}
}

// New returns a FilesCreatorService with all of the expected middleware wired in.
func New(middleware []Middleware) FilesCreatorService {
	var svc FilesCreatorService = NewBasicFilesCreatorService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
