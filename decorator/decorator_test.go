package decorator

import (
	"testing"
)

func TestDecorator(t *testing.T) {
	d1 := &stringDisplay{&baseDisplay{}, "A"}
	result := d1.show(d1)

	if result != "A" {
		t.Errorf("Expect result to equal %s, but %s.\n", "A", result)
	}

	d2 := &sideBorder{&border{display: d1}, "|"}
	result = d2.show(d2)

	if result != "|A|" {
		t.Errorf("Expect result to equal %s, but %s.\n", "|A|", result)
	}

}
