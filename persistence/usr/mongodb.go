package usr

import (
	"fmt"
	"log"

	mgo "gopkg.in/mgo.v2"
)

//SyncServDAO syncservice data access struct
type SyncServDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

var host []string

//COLLECTION mongo collection
const (
	COLLECTION = "syncserv"
	DATABASE   = "perf"
	SERVER     = "127.0.0.1:27017"
)

//Connect establish a connection to server
func Connect() *mgo.Database {
	session, err := mgo.Dial(SERVER)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(DATABASE)
	return db
}

//Store data to mongo
func (user Syncdata) Store() error {
	err := Connect().C(COLLECTION).Insert(user)
	fmt.Println("Data Inserted...")
	return err
}

func UserActions(r Repository) {
	r.Store()
}
