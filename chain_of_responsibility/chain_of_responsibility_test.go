package chain_of_responsibility

import (
	"testing"
)

func TestChainOfResponsibility(t *testing.T) {
	a := NewNoSupport("A")
	b := NewLimitSupport("B", 2)
	c := NewLimitSupport("C", 3)
	a.SetNext(b)
	b.SetNext(c)

	result := a.Handle(a, Trouble{1})
	expect := "trouble:1 is resolved by B"
	if result != expect {
		t.Errorf("Expect result to be equal %s, but %s\n", expect, result)
	}

	result = a.Handle(a, Trouble{2})
	expect = "trouble:2 is resolved by C"
	if result != expect {
		t.Errorf("Expect result to be equal %s, but %s\n", expect, result)
	}

	result = a.Handle(a, Trouble{3})
	expect = "trouble:3 cannot be resolved"
	if result != expect {
		t.Errorf("Expect result to be equal %s, but %s\n", expect, result)
	}
}
