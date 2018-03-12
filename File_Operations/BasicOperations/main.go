package main

import (
	"log"
	"os"
)

var (
	newfile *os.File
	err     error
)

func main() {
	newfile, err = os.Create("anupam.txt")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(newfile)
	newfile.Close()
}
