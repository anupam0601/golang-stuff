package main

import (
	"github.com/boltdb/bolt"
	"time"
	"fmt"
	"strconv"
)

func save(db *bolt.DB, workerNum string) error{
	err := db.Update(func(tx *bolt.Tx) error {
		dat, err := tx.CreateBucketIfNotExists([]byte("userAnupam"))
		if err != nil{
			return  err
		}
		return dat.Put([]byte(workerNum),[]byte("12:00:36"))
	})
	return err
}

func GetDataAll(db *bolt.DB, bucketName []byte) error{
	err := db.View(func(db *bolt.Tx) error {
		b := db.Bucket(bucketName)

		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			fmt.Printf("FileName=>[%s], TimeToSFTP=[%s]\n", k, v)
		}
		return nil
	})

	if err != nil {
		fmt.Println("failure : %s\n", err)
	}
	return nil
}

func main(){
	db, err := bolt.Open("perfStats.db", 0600, &bolt.Options{Timeout: 2 * time.Second})
	if err != nil {
		panic(err)
	}
	defer db.Close()
	workers := 10
	for i:=0; i<workers; i++ {
		go save(db, strconv.Itoa(i))
		//time.Sleep(1 *time.Second)
		fmt.Println(i)
	}
	time.Sleep(2 *time.Second)
	GetDataAll(db, []byte("userAnupam"))

}