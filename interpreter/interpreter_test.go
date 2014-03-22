package interpreter

import (
	"testing"
)

func TestInterpreter(t *testing.T) {

	program := "program go right end"

	node := ProgramNode{}
	context := NewContext(program)
	node.Parse(context)

	expect := "program: go right "
	if node.ToString() != expect {
		t.Errorf("Expect result to equal %s, but %s.\n", expect, node.ToString())
	}

}
