//File for experiments
package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

// total number of go routines to use at max
const maxgoroutines = 2

func createFiles(a string, chUpstream chan UpstreamChan){
	fmt.Println("Creating file :>>  ", a)
	f, err := os.Create(a)
	if err != nil {
		chUpstream <- UpstreamChan{"", err}
		//fmt.Println(err)
	}
	//chUpstream <- a
	chUpstream <- UpstreamChan{a, nil}
	defer f.Close()
	return
}

//Channel struct for upstream with data and error
type UpstreamChan struct {
	Output string
	Err error
}

func main(){
	var ch = make(chan string, 100) //It can be anything as long it is greater than maxgoroutines
	var chUpstream = make(chan UpstreamChan, 100) //Error and output channel
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
				createFiles(a, chUpstream)
				//time.Sleep(2 * time.Second)
			}
		}()
	}

	// Add the files to the channel , which is used as queue
	for i:=0; i<50; i++{
		t := strconv.Itoa(i)
		ch <- t //add i to queue
	}

	close(ch) //This tells the goroutines there is nothing else to do, channel is closed
	wg.Wait() //Wait for the goroutines to finish
	close(chUpstream)
	fmt.Println("Waiting.. 2 Secs....")
	time.Sleep(2 * time.Second)

	for r := range chUpstream{
		if r.Err != nil {
			fmt.Println(r.Err)
		}
		fmt.Println(r.Output)
		fmt.Println(r)
	}

	//select {
	//case msg1 := <-chUpstream:
	//	fmt.Println(msg1)
	//case msg2 := <-chUpstream:
	//	fmt.Println(msg2)
	//}
}