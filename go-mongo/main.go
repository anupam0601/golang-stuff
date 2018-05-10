package main

import (
	"fmt"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Game struct {
	Winner       string    `bson:"winner"`
	OfficialGame bool      `bson:"official_game"`
	Location     string    `bson:"location"`
	StartTime    time.Time `bson:"start"`
	EndTime      time.Time `bson:"end"`
	Players      []Player  `bson:"players"`
}

type Player struct {
	Name   string    `bson:"name"`
	Decks  [2]string `bson:"decks"`
	Points uint8     `bson:"points"`
	Place  uint8     `bson:"place"`
}

func NewPlayer(name, firstDeck, secondDeck string, points, place uint8) Player {
	return Player{
		Name:   name,
		Decks:  [2]string{firstDeck, secondDeck},
		Points: points,
		Place:  place,
	}
}

var isDropMe = true

func main() {
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

	game := Game{
		Winner:       "Dave",
		OfficialGame: true,
		Location:     "Austin",
		StartTime:    time.Date(2015, time.February, 12, 04, 11, 0, 0, time.UTC),
		EndTime:      time.Now(),
		Players: []Player{
			NewPlayer("Dave", "Wizards", "Steampunk", 21, 1),
			NewPlayer("Javier", "Zombies", "Ghosts", 18, 2),
			NewPlayer("George", "Aliens", "Dinosaurs", 17, 3),
			NewPlayer("Seth", "Spies", "Leprechauns", 10, 4),
		},
	}

	if isDropMe {
		err = session.DB("test").DropDatabase()
		if err != nil {
			panic(err)
		}
	}

	// Collection
	c := session.DB(Database).C(Collection)

	// Insert
	if err := c.Insert(game); err != nil {
		panic(err)
	}

	// Find and Count
	player := "Dave"
	gamesWon, err := c.Find(bson.M{"winner": player}).Count()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s has won %d games.\n", player, gamesWon)

}
