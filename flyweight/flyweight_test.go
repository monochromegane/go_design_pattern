package flyweight

import (
	"testing"
)

func TestFlyWeight(t *testing.T) {
	bigStr := NewBigString("121")
	result := bigStr.Print()

	expect := "-\n--\n-\n"
	if result != expect {
		t.Errorf("Expect result to equal %s, but %s.\n", expect, result)
	}
}
