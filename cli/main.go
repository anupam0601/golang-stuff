package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"time"

	sftpcl "github.com/anupam0601/golang-stuff/cli/awsftp"
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
func createGzFile() string {
	var b bytes.Buffer
	w := gzip.NewWriter(&b)
	w.Write([]byte("hello, world\n"))
	w.Close() // You must close this first to flush the bytes to the buffer.
	err := ioutil.WriteFile(path, b.Bytes(), 0666)
	if err != nil {
		panic(err)
	}
	fmt.Println("==> done creating gz file")
	return path
}

// Create zip file
// Taking path as the return value from createTxtFile func
func createZipFile() string {
	var pathZip = currentPath() + "/" + ts + os.Args[3]
	err := archiver.Zip.Make(pathZip, []string{createTxtFile()})
	if err != nil {
		panic(err)
	}
	fmt.Println("==> done creating Zip file")
	return pathZip
}

// Main Block
func main() {
	if os.Args[2] == "txt" {
		var txtFilePath = createTxtFile()
		sftpcl.Awsftp(txtFilePath)
	} else if os.Args[2] == "gz" {
		var gzFilePath = createGzFile()
		sftpcl.Awsftp(gzFilePath)
	} else if os.Args[2] == "zip" {
		var zipFilePath = createGzFile()
		sftpcl.Awsftp(zipFilePath)
	} else {
		log.Fatal("file extension is not valid")
	}
	// Command line commands execution
	out, err := exec.Command("ls").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The date is %s\n", out)
}
