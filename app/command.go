package app

import (
	"github.com/donutloop/go-blog-rest/config"
	"github.com/BurntSushi/toml"
	"log"
)

type command interface {
	Execute() interface{}
}

type MacroCommand struct {
	commands map[string]command
}

func (self *MacroCommand) Execute() map[string]interface{} {
	result := make(map[string]interface{})

	for index, command := range self.commands {
		result[index] = command.Execute()
	}

	self.Clear()

	return result
}

func (self *MacroCommand) Append(index string, command command) {
	self.commands[index] = command
}

func (self *MacroCommand) Clear() {
	self.commands = map[string]command{}
}

type LoadConfigurationCommand struct{
	ConfigFile string
}

func (self LoadConfigurationCommand) Execute() interface{} {

	var config config.Configuration

	if _, err := toml.DecodeFile(self.ConfigFile, &config); err != nil {
		log.Fatal(err)
	}

	return config
}
