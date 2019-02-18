package filescreatormicro

import (
	"encoding/json"
	"github.com/go-kit/kit/log"
	"time"
)

/// implement function to return ServiceMiddleware
func LoggingMiddleware(logger log.Logger) ServiceMiddleware {
	return func(next FilesCreatorService) FilesCreatorService {
		return loggingMiddleware{next, logger}
	}
}

// Make a new type and wrap into Service interface
// Add logger property to this type
type loggingMiddleware struct {
	FilesCreatorService
	logger log.Logger
}

// Implement Service Interface for LoggingMiddleware
func (mw loggingMiddleware) Create(fileDescriptor FileDescriptor) (output string, err error) {
	//Marshaling the incoming json from req body'FileDescriptor' struct
	//Otherwise logging will complain of unsupported type
	fileStr, _ := json.Marshal(fileDescriptor)
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "Create",
			"fileDescriptor", fileStr,
			"result", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	output, err = mw.FilesCreatorService.Create(fileDescriptor)
	return
}
