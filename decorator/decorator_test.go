package decorator

import (
	"testing"
)

func TestDecorator(t *testing.T) {
	d1 := NewStringDisplay("A")
	result := d1.Show(d1)

	if result != "A" {
		t.Errorf("Expect result to equal %s, but %s.\n", "A", result)
	}

	d2 := NewSideBorder(d1, "|")
	result = d2.Show(d2)

	if result != "|A|" {
		t.Errorf("Expect result to equal %s, but %s.\n", "|A|", result)
	}

}
