package filescreator

import (
	"context"
	"errors"
)

var (
	ErrFileNotCreated = errors.New("Coudn't create file")
)

// FilesCreator descriibes the Files create service
type FilesCreator interface {
	Create(ctx context.Context, fileDescriptor FileDescriptor) (string, error)
}
