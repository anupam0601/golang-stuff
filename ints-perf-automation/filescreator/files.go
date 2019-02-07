package filescreator

import "context"

// FileDescriptor represents file attributes
type FileDescriptor struct {
	ID            string `json:"id,omitempty"`
	Name          string `json:"name"`
	NumberOfFiles int32  `json:"number_of_files"`
	FileType      string `json:"file_type"`
}

type Repository interface {
	CreateFile(ctx context.Context, fileDescriptor FileDescriptor) error
}
