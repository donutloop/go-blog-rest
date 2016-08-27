package clog

import (
	"log"
	"fmt"
	"sync"
)

type cLogger struct {}

func New() *cLogger {
	return &cLogger{}
}

var instance *cLogger
var once sync.Once

// Singleton pattern
func GetInstance() *cLogger {

	once.Do(func() {
		instance = &cLogger{}
	})

	return instance
}

func (self cLogger) Warning(v map[string]interface{}) {
	message := self.formatLog("Warning",v)
	log.Print(message)
}

func (self cLogger) Info(v map[string]interface{}) {
	message := self.formatLog("Info", v)
	log.Print(message)
}

func (self cLogger) Fatal(v map[string]interface{}) {
	message := self.formatLog("Fatal error", v)
	log.Print(message)
}

func (self cLogger) Debug(v map[string]interface{}) {
	message := self.formatLog("Debug", v)
	log.Fatal(message)
}

func (self cLogger) formatLog(prefix string , v map[string]interface{}) string{

	log.SetPrefix("["+ prefix +"][")

	message := ""

	for index, value := range v {
		message += fmt.Sprintf("[%s : %s]", index, value)
	}

	return "]" + message
}