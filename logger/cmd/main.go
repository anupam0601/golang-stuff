package main

import (
	li "github.com/anupam0601/golang-stuff/logger"
)

func main() {

	var standardLogger = li.NewLogger()

	// You can then call a method of our standard logger in the context of an error
	// you would like to log.

	standardLogger.AttributeNotPresent("Anupam")
	standardLogger.Step("1", "Upload file to SFTP")

}