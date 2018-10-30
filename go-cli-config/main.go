package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// Config sync-service config
type Config struct {
	Username string `json:"username"`
	Password string `json:"password"`
	S3path   string `json:"s3path"`
	Region   string `json:"region"`
}

// Take user input and create config json
func (c *Config) takeinput() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Username: ")
	text, _ := reader.ReadString('\n')
	//trim new line
	text = strings.TrimSuffix(text, "\n")
	c.Username = text

	fmt.Println("Enter Password: ")
	text2 := ""
	fmt.Scanln(&text2)
	// fmt.Println(&text2)
	c.Password = text2

	//Marshal json and write to a file
	configJSON, _ := json.Marshal(c)
	fmt.Println(string(configJSON))
	jsonFile, err := os.Create("config.json")
	if err != nil {
		panic(err)
	}
	jsonFile.Write(configJSON)

}

func main() {
	c := Config{} //Empty intialization
	c.takeinput()
}
