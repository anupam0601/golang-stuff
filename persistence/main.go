package main

import "github.com/anupam0601/golang-stuff/persistence/usr"

func main() {
	user := usr.Syncdata{S3path: "Geetanjali", Username: "Ananya", Password: "Debnath"}
	usr.UserActions(user)
}
