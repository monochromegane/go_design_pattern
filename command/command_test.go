package command

import (
	"testing"
)

func TestCommand(t *testing.T) {

	macro := macroCommand{}

	macro.append(&drawCommand{&position{1, 1}})
	macro.append(&drawCommand{&position{2, 2}})

	expect := "1.1\n2.2\n"
	if macro.execute() != expect {
		t.Errorf("Expect result to equal %s, but %s.\n", expect, macro.execute())
	}

	macro.undo()
	expect = "1.1\n"
	if macro.execute() != expect {
		t.Errorf("Expect result to equal %s, but %s.\n", expect, macro.execute())
	}

}
