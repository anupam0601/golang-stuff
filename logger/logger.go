package logger

import (
	"github.com/sirupsen/logrus"
)

// Event stores messages to log later, from our standard interface
type Event struct {
	id      int
	message string
}

// StandardLogger enforces specific log message formats
type StandardLogger struct {
	*logrus.Logger
}

// NewLogger initializes the standard logger
func NewLogger() *StandardLogger {
	var baseLogger = logrus.New()
	var standardLogger = &StandardLogger{baseLogger}
	standardLogger.Formatter = &logrus.JSONFormatter{}
	return standardLogger
}

// Declare variables to store log messages as new Events
var (
	attributeNotPresent = Event{id: 3000, message: "Attribute not present: %s"}
	step = Event{message:"STEP %s: %s"}
)

// AttributeNotPresent is a standard error message
func (l *StandardLogger) AttributeNotPresent(attrName string) {
	l.Errorf(attributeNotPresent.message, attrName)
}

// Step is an info level message to log test steps
func(l * StandardLogger) Step(stepNumb string, stepMsg string){
	//l.Info(attrName)
	l.Infof(step.message, stepNumb, stepMsg)
}
