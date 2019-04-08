package main

import "fmt"

type filesService interface {
	upload(endpoint string) string
}

type fileMeta struct {
	FileName string
	FileType string
}

// Implementing interface method with struct
func (a *fileMeta) upload(endpoint string) string {
	fmt.Println("Uploading to endpoint", endpoint)
	return a.FileName + "--" + a.FileType
}

// Method implements the interface by returning method of interface
func uploadToEndpoint(a filesService) string {
	return a.upload("SFTP")
}

func main() {
	funcInit := uploadToEndpoint(&fileMeta{FileName: "Anupam.txt", FileType: "text"})
	fmt.Println(funcInit)
}
