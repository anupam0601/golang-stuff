package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"
)

// Path to create file
var path = currentPath() + "/" + os.Args[1]

// Unix Time stamp
var t = time.Now().Unix()

// Get current path
func currentPath() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	return exPath
}

// Create txt file
func createTxtFile() {
	// open file using READ & WRITE permission
	data := []byte("hello world\n")
	err := ioutil.WriteFile(path, data, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("==> done writing to file")
}

// Create gz file
func createGzFile() {
	var b bytes.Buffer
	w := gzip.NewWriter(&b)
	w.Write([]byte("hello, world\n"))
	w.Close() // You must close this first to flush the bytes to the buffer.
	err := ioutil.WriteFile(path, b.Bytes(), 0666)
	if err != nil {
		panic(err)
	}
	fmt.Println("==> done creating gz file")
}

func main() {
	if os.Args[2] == "txt" {
		createTxtFile()
	} else if os.Args[2] == "gz" {
		createGzFile()
	} else {
		log.Fatal("file extension is not valid")
	}

	log.Println(t)

}
