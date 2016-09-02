package command

import (
	"errors"
)

type CommandWrapper struct {
	Name string
	Command command
}

type command interface {
	Execute() interface{}
}

// Command pattern
type MacroCommand struct {
	CommandWrappers []CommandWrapper
}

// Execute all append commands (First input first output)
func (self *MacroCommand) Execute() (map[string]interface{}, error){
	result := make(map[string]interface{})

	for _, commandWrapper := range self.CommandWrappers {

		if commandWrapper.Name == "" || commandWrapper.Command == nil {
			return nil, errors.New("Command is not properly configured")
		}

		result[commandWrapper.Name] = commandWrapper.Command.Execute()
	}

	self.clear()

	return result, nil
}

//appends commandWrapper elements to the end of a the CommandWrappers slice.
func (self *MacroCommand) Append(commandWrapper CommandWrapper) {
	self.CommandWrappers = append(self.CommandWrappers, commandWrapper)
}

// Overwrite the current CommandWrappers slice with a empty slice
func (self *MacroCommand) clear() {
	self.CommandWrappers = []CommandWrapper{}
}
