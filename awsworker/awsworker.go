package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "awsWorker"
	app.Usage = "Does common integration services jobs"
	app.Action = func(c *cli.Context) error {
		fmt.Println("Welcome to aws Worker")
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
