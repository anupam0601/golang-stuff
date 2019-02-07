package main

import "log"

type FilesCreatorService interface {
	StoreFiles(FileType string, NumberOfFiles int)
}

type FilesCreator struct{}

func (c FilesCreator) StoreFiles(FileType string, NumberOfFiles int) {
	if FileType == "txt" && NumberOfFiles == 10 {
		log.Println("Creating text files of type and quantity :", FileType, NumberOfFiles)
	}
}

func UserCreateFiles(s FilesCreatorService) {
	s.StoreFiles("txt", 10)
}
func main() {
	r := FilesCreator{}
	UserCreateFiles(r)
}
