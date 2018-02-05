package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	data, err := ioutil.ReadFile("file_to_read.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}
