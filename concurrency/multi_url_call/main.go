package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	mgo "gopkg.in/mgo.v2"
)

//HTTPResponse struct
type HTTPResponse struct {
	status string
	body   []byte
}

//DB struct
type datmg struct {
	Stat string   `bson:"status"`
	Dat  []string `bson:"data"`
}

func main() {
	//Define a new channel
	var ch chan HTTPResponse = make(chan HTTPResponse)

	// List of APIs to call
	urls := [2]string{"https://jsonplaceholder.typicode.com/posts/1", "https://jsonplaceholder.typicode.com/posts/1/comments"}
	for _, url := range urls {
		//For each URL call the DOHTTPGet function (notice the go keyword)
		// Sender
		go DoHTTPGet(url, ch)

	}

	// Getting data in a var. Receiver (Blocking)
	// Running for loop on channel data and storing it
	// in a variable
	for n := range ch {
		// fmt.Println(n.status)
		newmgdat := datmg{
			Stat: n.status,
			Dat:  []string{string(n.body)},
		}

		fmt.Println(newmgdat)

		// Mongo DB Operations follows
		Host := []string{
			"127.0.0.1:27017",
			// replica set addrs...
		}
		const (
			Database   = "Godb"
			Collection = "GoColl"
		)
		session, err := mgo.DialWithInfo(&mgo.DialInfo{
			Addrs: Host,
			// Username: Username,
			// Password: Password,
			// Database: Database,
			// DialServer: func(addr *mgo.ServerAddr) (net.Conn, error) {
			// 	return tls.Dial("tcp", addr.String(), &tls.Config{})
			// },
		})
		if err != nil {
			panic(err)
		}
		defer session.Close()

		// Collection
		c := session.DB(Database).C(Collection)

		// Insert
		if err := c.Insert(newmgdat); err != nil {
			panic(err)
		}
	}
}

//DoHTTPGet : HTTP get function
func DoHTTPGet(url string, ch chan<- HTTPResponse) {
	//Execute the HTTP get
	httpResponse, _ := http.Get(url)
	httpBody, _ := ioutil.ReadAll(httpResponse.Body)
	//Send an HTTPResponse back to the channel
	ch <- HTTPResponse{httpResponse.Status, httpBody}
}
