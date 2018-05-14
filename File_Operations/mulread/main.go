package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"sync"
)

var wg sync.WaitGroup

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

func main() {
	coll := make(chan string)

	go func() {
		scanner := bufio.NewScanner("file1.txt")
		for scanner.Scan() {
			coll <- scanner.Text()
		}
		close(coll)
	}()

	// coll <- data

	// go func() {
	// 	datone := readOne("file2.txt")
	// 	coll <- datone
	// }()

	// go func() {
	// 	for {
	// 		fmt.Println(<-coll)
	// 	}
	// }()
}
