package main

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"os"
)

const (
	chunksize int = 32
)

var (
	data  *os.File
	part  []byte
	err   error
	count int
)

func openFile(name string) (byteCount int, buffer *bytes.Buffer) {
	data, err = os.Open(name)
	if err != nil {
		log.Println(err)
	}
	defer data.Close()

	reader := bufio.NewReader(data)
	buffer = bytes.NewBuffer(make([]byte, 0))
	part = make([]byte, chunksize)

	for {
		if count, err = reader.Read(part); err != nil {
			break
		}
		buffer.Write(part[:count])
	}
	if err != io.EOF {
		log.Fatal("Error Reading ", name, ": ", err)
	} else {
		err = nil
	}
	byteCount = buffer.Len()
	return
}

func main() {
	//b := bytes.NewBuffer(make([]byte, 64))
	//proverbs := new(bytes.Buffer)
	//proverbs.WriteString("Channels orchestrate mutexes serialize\n")
	//proverbs.WriteString("Cgo is not Go\n")
	//proverbs.WriteString("Errors are values\n")
	//proverbs.WriteString("Don't panic\n")
	_, partData := openFile("data.txt")
	pr, pw := io.Pipe()

	go func() {
		defer pw.Close()
		_, _ = io.Copy(pw, partData)
	}()

	_, _ = io.Copy(os.Stdout, pr)
	pr.Close()
}
