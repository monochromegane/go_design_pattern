package chain_of_responsibility

import (
	"testing"
)

func TestChainOfResponsibility(t *testing.T) {
	a := newNoSupport("A")
	b := newLimitSupport("B", 2)
	c := newLimitSupport("C", 3)
	a.setNext(b)
	b.setNext(c)

	result := a.handle(a, trouble{1})
	expect := "trouble:1 is resolved by B"
	if result != expect {
		t.Errorf("Expect result to be equal %s, but %s\n", expect, result)
	}

	result = a.handle(a, trouble{2})
	expect = "trouble:2 is resolved by C"
	if result != expect {
		t.Errorf("Expect result to be equal %s, but %s\n", expect, result)
	}

	result = a.handle(a, trouble{3})
	expect = "trouble:3 cannot be resolved"
	if result != expect {
		t.Errorf("Expect result to be equal %s, but %s\n", expect, result)
	}
}
