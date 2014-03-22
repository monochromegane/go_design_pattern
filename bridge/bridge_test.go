package bridge

import (
	"testing"
)

func TestBridge(t *testing.T) {

	d1 := DefaultDisplay{&StringDisplayImpl{"AAA"}}
	expect := "+---+\n|AAA|\n+---+\n"
	if result := d1.Display(); result != expect {
		t.Errorf("Expect result to equal %s, but %s.\n", expect, result)
	}

	d2 := CountDisplay{&DefaultDisplay{&StringDisplayImpl{"BBB"}}}
	expect = "+---+\n|BBB|\n+---+\n"
	if result := d2.Display(); result != expect {
		t.Errorf("Expect result to equal %s, but %s.\n", expect, result)
	}

	expect = "+---+\n|BBB|\n|BBB|\n+---+\n"
	if result := d2.MultiDisplay(2); result != expect {
		t.Errorf("Expect result to equal %s, but %s.\n", expect, result)
	}
}
