package service

import (
	"context"

	log "github.com/go-kit/kit/log"
)

type Middleware func(FilesCreatorService) FilesCreatorService

type loggingMiddleware struct {
	logger log.Logger
	next   FilesCreatorService
}

func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next FilesCreatorService) FilesCreatorService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) Create(ctx context.Context, fileDescriptor FileDescriptor) (s0 string) {
	defer func() {
		l.logger.Log("method", "Create", "fileDescriptor", fileDescriptor, "s0", s0)
	}()
	return l.next.Create(ctx, fileDescriptor)
}
