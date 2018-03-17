package main

import (
	"log"
	"net/smtp"
)

func init() {
	log.SetPrefix("LOGSYSTEM: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	log.Println("init started")
}

func main() {
	// Connect to smtp server
	client, err := smtp.Dial("smtp.smail.com:25")
	if err != nil {
		log.Fatalln(err)
	}
	client.Data()
}
