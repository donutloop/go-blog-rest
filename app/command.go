package app

import (
	"github.com/donutloop/go-blog-rest/config"
	"github.com/BurntSushi/toml"
	"log"
)

type command interface {
	Execute() interface{}
}

func newCommandChain() *MacroCommand{
	return &MacroCommand{commands:map[string]command{"config":LoadConfigurationCommand{}}}
}

type MacroCommand struct {
	commands map[string]command
}

func (self *MacroCommand) Execute() map[string]interface{} {
	result := make(map[string]interface{})

	for index, command := range self.commands {
		result[index] = command.Execute()
	}

	return result
}

func (self *MacroCommand) Append(index string, command command) {
	self.commands[index] = command
}

func (self *MacroCommand) Clear() {
	self.commands = map[string]command{}
}

const CONFIGURATION_FILE string = "./config/config.toml"

type LoadConfigurationCommand struct{}

func (self LoadConfigurationCommand) Execute() interface{} {

	var config config.Configuration

	if _, err := toml.DecodeFile(CONFIGURATION_FILE, &config); err != nil {
		log.Fatal(err)
	}

	return config
}