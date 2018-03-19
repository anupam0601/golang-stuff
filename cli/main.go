package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/mholt/archiver"
)

// Unix Time stamp converted to string
var ts = strconv.FormatInt(time.Now().Unix(), 10)

// Path to create file
var path = currentPath() + "/" + ts + os.Args[1]

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
func createTxtFile() string {
	// open file using READ & WRITE permission
	data := []byte("hello world\n")
	err := ioutil.WriteFile(path, data, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("==> done writing to file")
	fmt.Println(path)
	return path
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

// Create zip file
// Taking path as the return value from createTxtFile func
func createZipFile() {
	var pathZip = currentPath() + "/" + os.Args[3]
	err := archiver.Zip.Make(pathZip, []string{createTxtFile()})
	if err != nil {
		panic(err)
	}
	fmt.Println("==> done creating Zip file")
}

// Main Block
func main() {
	if os.Args[2] == "txt" {
		createTxtFile()
	} else if os.Args[2] == "gz" {
		createGzFile()
	} else if os.Args[2] == "zip" {
		createZipFile()
	} else {
		log.Fatal("file extension is not valid")
	}

	// log.Println(ts)
}
