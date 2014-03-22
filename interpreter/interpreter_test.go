package interpreter

import (
	"testing"
)

func TestInterpreter(t *testing.T) {

	program := "program go right end"

	node := programNode{}
	context := NewContext(program)
	node.parse(context)

	expect := "program: go right "
	if node.toString() != expect {
		t.Errorf("Expect result to equal %s, but %s.\n", expect, node.toString())
	}

}
