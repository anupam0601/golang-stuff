package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type HTTPResponse struct {
	status string
	body   []byte
}

func main() {
	//Define a new channel
	var ch chan HTTPResponse = make(chan HTTPResponse)

	// List of APIs to call
	urls := [2]string{"https://jsonplaceholder.typicode.com/posts/1", "https://jsonplaceholder.typicode.com/posts/1/comments"}
	for _, url := range urls {
		//For each URL call the DOHTTPGet function (notice the go keyword)
		go DoHTTPGet(url, ch)

	}
	// Get the response
	for range urls {
		fmt.Println((<-ch).status)
		// fmt.Println((<-ch).body)
	}
}
func DoHTTPGet(url string, ch chan<- HTTPResponse) {
	//Execute the HTTP get
	httpResponse, _ := http.Get(url)
	httpBody, _ := ioutil.ReadAll(httpResponse.Body)
	//Send an HTTPResponse back to the channel
	ch <- HTTPResponse{httpResponse.Status, httpBody}
}
