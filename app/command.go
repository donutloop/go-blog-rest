package app

import (
	"github.com/donutloop/go-blog-rest/config"
	"github.com/BurntSushi/toml"
	"log"
	"github.com/donutloop/go-blog-rest/utils/command"
)

func newCommandChain() *command.MacroCommand {

	commands := []command.CommandWrapper{
		command.CommandWrapper{Name:"config",Command:LoadConfigurationCommand{ConfigFile:CONFIGURATION_FILE}},
	}

	return &command.MacroCommand{CommandWrappers:commands}
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
