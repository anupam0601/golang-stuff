package service

import (
	"context"
	"log"
	"strconv"

	"github.com/gofrs/uuid"
)

// FilesCreatorService describes the service.
type FilesCreatorService interface {
	// Add your methods here
	// e.x: Foo(ctx context.Context,s string)(rs string, err error)
	Create(ctx context.Context, fileDescriptor FileDescriptor) string
}

type FileDescriptor struct {
	ID            string `json:"id,omitempty"`
	Name          string `json:"name"`
	NumberOfFiles string `json:"number_of_files"`
	FileType      string `json:"file_type"`
}

type basicFilesCreatorService struct{}

func (b *basicFilesCreatorService) Create(ctx context.Context, fileDescriptor FileDescriptor) (s0 string) {
	// TODO implement the business logic of Create

	totalfilesToCreate, err := strconv.Atoi(fileDescriptor.NumberOfFiles)
	if err != nil {
		log.Println(err)
	}

	// FileName := id + "-" + fileDescriptor.Name

	if fileDescriptor.FileType == "txt" {
		for i := 0; i <= totalfilesToCreate; i++ {
			go func() {
				uuid, _ := uuid.NewV4()
				id := uuid.String()
				fileDescriptor.ID = id
				log.Println(id + "-" + fileDescriptor.Name)
				// log.Println("Unique File ====>", FileName)
			}()
		}
	}
	return fileDescriptor.Name
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
