//File for experiments
package main

import (
	"fmt"
	"sync"
)

// total number of go routines to use at max
const maxgoroutines = 10

func createFiles(a int){
	fmt.Println("Creating file :>>  ", a)
	return
}

func main(){
	var ch = make(chan int, 50) //It can be anything as long it is greater than maxgoroutines
	var wg sync.WaitGroup

	//This block starts maxgoroutines that wait for executing some job
	wg.Add(maxgoroutines)
	for i:=0; i<maxgoroutines; i++{
		go func() {
			fmt.Println("Spawning go routine : >>>> ", i)
			for {
				a,ok := <-ch
				if !ok {
					wg.Done() // if there is no file to create and the channel has been closed then end the goroutine
					return
				}
				createFiles(a)
			}
		}()
	}

	// Add the files to the channel , which is used as queue
	for i:=0; i<50; i++{
		ch <- i //add i to queue
	}

	close(ch) //This tells the goroutines there is nothing else to do, channel is closed
	wg.Wait() //Wait for the goroutines to finish
}