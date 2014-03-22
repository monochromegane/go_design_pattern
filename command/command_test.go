package command

import (
	"testing"
)

func TestCommand(t *testing.T) {

	macro := MacroCommand{}

	macro.Append(&DrawCommand{&Position{1, 1}})
	macro.Append(&DrawCommand{&Position{2, 2}})

	expect := "1.1\n2.2\n"
	if macro.Execute() != expect {
		t.Errorf("Expect result to equal %s, but %s.\n", expect, macro.Execute())
	}

	macro.Undo()
	expect = "1.1\n"
	if macro.Execute() != expect {
		t.Errorf("Expect result to equal %s, but %s.\n", expect, macro.Execute())
	}

}
