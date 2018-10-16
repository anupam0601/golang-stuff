package main

import "fmt"

func fanout(In <-chan int, OutA, OutB chan int) {
	for data := range In { // Receive data until closed
		select { // Send to first non blocking channel
		case OutA <- data:
			fmt.Println(<-OutA)
		case OutB <- data:
			fmt.Println(<-OutA)
			// fmt.Println(data)
		}

	}
}

func main() {
	In := make(chan int)
	OutA := make(chan int)
	OutB := make(chan int)
	go func() {
		In <- 2
		In <- 32
		close(In)
	}()
	go fanout(In, OutA, OutB)
}
