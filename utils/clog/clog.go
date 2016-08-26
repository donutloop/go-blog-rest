package clog

import (
	"log"
	"fmt"
)

type cLogger struct {}

func New() *cLogger {
	return &cLogger{}
}

func (self *cLogger) Warning(v map[string]interface{}) {
	message := self.formatLog("Warning",v)
	log.Print(message)
}

func (self *cLogger) Info(v map[string]interface{}) {
	message := self.formatLog("Info", v)
	log.Print(message)
}

func (self *cLogger) Fatal(v map[string]interface{}) {
	message := self.formatLog("Fatal error", v)
	log.Print(message)
}

func (self *cLogger) Debug(v map[string]interface{}) {
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