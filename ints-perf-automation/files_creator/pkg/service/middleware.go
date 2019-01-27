package service

import (
	"context"

	log "github.com/go-kit/kit/log"
)

// Middleware describes a service middleware.
type Middleware func(FilesCreatorService) FilesCreatorService

type loggingMiddleware struct {
	logger log.Logger
	next   FilesCreatorService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a FilesCreatorService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next FilesCreatorService) FilesCreatorService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) CreateFiles(ctx context.Context, numberOfFiles int) (e0 error) {
	defer func() {
		l.logger.Log("method", "CreateFiles", "numberOfFiles", numberOfFiles, "e0", e0)
	}()
	return l.next.CreateFiles(ctx, numberOfFiles)
}
func (l loggingMiddleware) StoreFileMetadata(ctx context.Context, fileName string, timeStamp string) (e0 error) {
	defer func() {
		l.logger.Log("method", "StoreFileMetadata", "fileName", fileName, "timeStamp", timeStamp, "e0", e0)
	}()
	return l.next.StoreFileMetadata(ctx, fileName, timeStamp)
}
