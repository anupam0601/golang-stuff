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
	fmt.Println("Text file ===>", path)
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
	fmt.Println("Gz file name ==>", path)
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
	fmt.Println("Zip file name ==>", pathZip)
	return pathZip
}

func runCmd() {
	cmd := exec.Command("awslogs", "get", "sync-service", "--watch")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// Main Block
func main() {
	// Using switch case to identify file masks from command line args
	switch {
	case os.Args[2] == "txt":
		var txtFilePath = createTxtFile()
		sftpcl.Awsftp(txtFilePath)
	case os.Args[2] == "gz":
		var gzFilePath = createGzFile()
		sftpcl.Awsftp(gzFilePath)
	case os.Args[2] == "zip":
		var zipFilePath = createZipFile()
		sftpcl.Awsftp(zipFilePath)
	default:
		log.Fatal("Extension not valid, Please provide a proper extension")
	}
<<<<<<< 2d6e77c5f3985b584e246aab0940fed62bcd575c
	// Command line commands execution
	out, err := exec.Command("ls").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The date is %s\n", out)
=======
	runCmd()
>>>>>>> added cli features
}
