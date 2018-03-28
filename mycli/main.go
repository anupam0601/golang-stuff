package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/urfave/cli"
)

// Unix Time stamp converted to string
var ts = strconv.FormatInt(time.Now().Unix(), 10)

// Get current path
func currentPath() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	return exPath
}

// Create txt file
func createTxtFile(path string) string {
	// open file using READ & WRITE permission
	data := []byte("hello world\n")
	err := ioutil.WriteFile(path, data, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("==> done writing to file")
	fmt.Println(path)
	return path
}

func main() {
	app := cli.NewApp()

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "task",
			Usage: "Create a task, ta",
			Value: "createfile",
		},
		cli.StringFlag{
			Name:  "type, t",
			Usage: "type of file to create",
			Value: "txt",
		},
		cli.StringFlag{
			Name:  "filename, fn",
			Usage: "filename to create",
			Value: "anupam.txt",
		},
	}

	// With flags
	app.Action = func(c *cli.Context) error {
		// name := "acli"
		// if c.NArg() > 0 {
		// 	name = c.Args().Get(0)
		// }
		if c.String("task") == "createfile" && c.String("type") == "txt" && c.String("filename") == "anupam.txt" {
			fmt.Println("creating txt...")
		} else if c.String("task") == "createfile" && c.String("type") == "zip" && c.String("filename") == "anupam.zip" {
			fmt.Println("creating zip...")
		}
		return nil
	}

	// Without flags
	app.Commands = []cli.Command{
		{
			Name: "createfile",
			Subcommands: []cli.Command{
				{
					Name:  "txt",
					Usage: "create txt file",
					Action: func(c *cli.Context) error {
						txtPath := currentPath() + "/" + ts + c.Args().First()
						fmt.Println("Creating text file: ", c.Args().First())
						txtfile := createTxtFile(txtPath)
						fmt.Println("Done creating txt file ===>", txtfile)
						return nil
					},
				},
				{
					Name:  "gz",
					Usage: "Create gz file",
					Action: func(c *cli.Context) error {
						fmt.Println("Creating gz file: ", c.Args().First())
						return nil
					},
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
