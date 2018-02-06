package main

import (
	"fmt"
	"io/ioutil"
	"sync"
)

var wg sync.WaitGroup

func main() {
	// Concurrently reading two files and printing them using goroutine and wait group
	wg.Add(2)
	go readOne("file_1.txt")
	go readOne("file_2.txt")
	wg.Wait()

	// For sequential reading try below procedure
	readOne("file_1.txt")
	readOne("file_2.txt")
}

// function to read a file
func readOne(fileToRead string) {
	for i := 0; i < 40; i++ {
		data, err := ioutil.ReadFile(fileToRead)
		if err != nil {
			panic(err)
		}
		fmt.Println(fileToRead, string(data))
	}
	wg.Done()
}
