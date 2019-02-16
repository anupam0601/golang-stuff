package filescreatormicro

import "fmt"

type FilesCreatorService interface {
	Create(fileDescriptor FileDescriptor) (string, error)
}

// FileDescriptor file descriptor struct
type FileDescriptor struct {
	ID string `json:"id,omitempty"`
	Name string `json:"name"`
	FilesCount string `json:"files_count"`
	FileType string `json:"file_type"`
}


type BasicFilesCreator struct{}

// Implement
func (BasicFilesCreator) Create(fileDescriptor FileDescriptor) (string, error) {
	fmt.Println("File metadata ====>",fileDescriptor)
	return fileDescriptor.ID, nil
}
