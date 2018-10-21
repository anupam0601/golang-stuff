package main

import "fmt"

func main() {

	// For our example we'll select across two channels.
	c1 := make(chan int)
	// c2 := make(chan string)

	// Each channel will receive a value after some amount
	// of time, to simulate e.g. blocking RPC operations
	// executing in concurrent goroutines.
	go func() {
		// time.Sleep(1 * time.Second)
		for i := 0; i < 100; i++ {
			// c1 <- "one"
			// c1 <- "two"
			c1 <- i
		}
	}()
	// go func() {
	// 	time.Sleep(2 * time.Second)
	// 	c2 <- "two"
	// }()

	// We'll use `select` to await both of these values
	// simultaneously, printing each one as it arrives.
	// msg1 := make(chan string)
	// msg2 := make(chan string)

	// Load balanced between two cases
	for i := 0; i < 100; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received from msg1", msg1)
		case msg2 := <-c1:
			fmt.Println("received from msg2", msg2)
		}
	}
}
