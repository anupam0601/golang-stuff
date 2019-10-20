package main

import (
	"fmt"
	"github.com/anupam0601/golang-stuff/logger"
	"log"
)

func testLogger() string {
	return "Anupam"
}
func main() {
	config := logger.Configuration{
		EnableConsole:     true,
		ConsoleLevel:      logger.Debug,
		ConsoleJSONFormat: true,
	}
	err := logger.NewLogger(config, logger.InstanceZapLogger)
	if err != nil {
		log.Fatalf("Could not instantiate log %s", err.Error())
	}

	contextLogger := logger.WithFields(logger.Fields{"key1": "value1"})
	v:= testLogger()
	fmt.Println(v)
	logger.Infof(v)
	contextLogger.Infof(v)

	//contextLogger.Debugf("Starting with zap")
	//contextLogger.Infof("Zap is awesome")

}
