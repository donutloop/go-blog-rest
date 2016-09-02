package command

import "testing"

type testCommand struct{}

func (self testCommand) Execute() interface{}{
	return "Test-Echo"
}

func TestMacroCommand_Execute(t *testing.T) {
	commandWrapper := CommandWrapper{Name:"Test-Command", Command:testCommand{}}
	command := MacroCommand{}
	command.Append(commandWrapper)

	data, _ := command.Execute()

	if data["Test-Command"].(string) != "Test-Echo"{
		t.Error("Value isn't correct")
	}
}

func TestMacroCommand_Execute_Command_Name_Is_Missing(t *testing.T) {
	commandWrapper := CommandWrapper{Name:"", Command:testCommand{}}
	command := MacroCommand{}
	command.Append(commandWrapper)

	_, err := command.Execute()

	if err == nil{
		t.Error("Value isn't correct")
	}
}

func TestMacroCommand_Execute_Command_Is_Missing(t *testing.T) {
	commandWrapper := CommandWrapper{Name:"Test-Command", Command:nil}
	command := MacroCommand{}
	command.Append(commandWrapper)

	_, err := command.Execute()

	if err == nil{
		t.Error("Value isn't correct")
	}
}