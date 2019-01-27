// Business logic
package service

import "context"

// FilesCreatorService describes the service.
type FilesCreatorService interface {
	// Add your methods here
	CreateFiles(ctx context.Context, numberOfFiles int) error
	StoreFileMetadata(ctx context.Context, fileName string, timeStamp string) error
	// CreateFiles(typeOfFile string, numberOfFiles int) error
	// StoreFileMetadata(fileName string, timeStamp string) error
}

type basicFilesCreatorService struct{}

func (b *basicFilesCreatorService) CreateFiles(ctx context.Context, numberOfFiles int) (e0 error) {
	// TODO implement the business logic of CreateFiles
	return e0
}
func (b *basicFilesCreatorService) StoreFileMetadata(ctx context.Context, fileName string, timeStamp string) (e0 error) {
	// TODO implement the business logic of StoreFileMetadata
	return e0
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
