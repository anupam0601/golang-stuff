package main

import (
	"bufio"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// d1 := []byte("Hello Go...")
	// ioutil.WriteFile("filetowrite.txt", d1, 0644)
	f, err := os.Create("file.txt")
	check(err)
	defer f.Close()
	w := bufio.NewWriter(f)
	w.WriteString("Writing")
	w.Flush()

}
