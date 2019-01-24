// Business logic
package service

// FilesCreatorService describes the service.
type FilesCreatorService interface {
	// Add your methods here
	// e.x: Foo(ctx context.Context,s string)(rs string, err error)
	createFiles(typeOfFile string, numberOfFiles int)
	storeFileMetadata(fileName string, timeStamp string)
}
