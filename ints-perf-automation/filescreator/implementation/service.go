package implementation

import (
	"context"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/gofrs/uuid"

	filesvc "github.com/anupam0601/golang-stuff/ints-perf-automation/filescreator"
)

// service implements the FilesCreator service
type service struct {
	repository filesvc.Repository
	logger     log.Logger
}

// NewService creates and returns a new FilesCreator service instance
func NewService(rep filesvc.Repository, logger log.Logger) filesvc.FilesCreator {
	return &service{
		repository: rep,
		logger:     logger,
	}
}

func (s *service) Create(ctx context.Context, fileDescriptor filesvc.FileDescriptor) (string, error) {
	logger := log.With(s.logger, "method", "Create")
	uuid, _ := uuid.NewV4()
	id := uuid.String()
	fileDescriptor.ID = id
	fileDescriptor.FileType = "txt"
	fileDescriptor.Name = fileDescriptor.ID + "--- TEST"
	fileDescriptor.NumberOfFiles = 12

	if err := s.repository.CreateFile(ctx, fileDescriptor); err != nil {
		level.Error(logger).Log("err", err)
		return "", filesvc.ErrFileNotCreated
	}
	return fileDescriptor.Name, nil
}
