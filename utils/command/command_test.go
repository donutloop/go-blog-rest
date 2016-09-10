package command

import "testing"

type testCommand struct{}

func (self testCommand) Execute() interface{}{
	return "Test-Echo"
}

func TestMacroCommand_Execute(t *testing.T) {

	t.Log("create instance of a command wrapper")
	commandWrapper := CommandWrapper{Name:"Test-Command", Command:testCommand{}}

	t.Log("create a instance of a macro command")
	command := MacroCommand{}

	t.Log("push the command wrapper on the macro command stack")
	command.Append(commandWrapper)

	t.Log("call the method Execute of the Macro command")
	data, _ := command.Execute()

	t.Log("check if the return value is vailid")
	if data["Test-Command"].(string) != "Test-Echo"{
		t.Error("Value isn't correct")
	}
}

func TestMacroCommand_Execute_Command_Name_Is_Missing(t *testing.T) {

	t.Log("create instance of a command wrapper with bad parameters")
	commandWrapper := CommandWrapper{Name:"", Command:testCommand{}}

	t.Log("create a instance of a macro command")
	command := MacroCommand{}

	t.Log("push the command wrapper on the macro command stack")
	command.Append(commandWrapper)

	t.Log("call the method Execute of the Macro command")
	_, err := command.Execute()

	t.Log("check if an error is happened")
	if err == nil{
		t.Error("Value isn't correct")
	}
}

func TestMacroCommand_Execute_Command_Is_Missing(t *testing.T) {

	t.Log("create instance of a command wrapper with bad parameters")
	commandWrapper := CommandWrapper{Name:"Test-Command", Command:nil}

	t.Log("create a instance of a macro command")
	command := MacroCommand{}

	t.Log("push the command wrapper on the macro command stack")
	command.Append(commandWrapper)

	t.Log("call the method Execute of the Macro command")
	_, err := command.Execute()

	t.Log("check if an error is happened")
	if err == nil{
		t.Error("Value isn't correct")
	}
}

func BenchmarkMacroCommand(b *testing.B) {

	for n := 0; n < b.N; n++ {
		commandWrapper := CommandWrapper{Name:"Test-Command", Command:testCommand{}}
		command := MacroCommand{}
		command.Append(commandWrapper)
		command.Execute()
	}
}