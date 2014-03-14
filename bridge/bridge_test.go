package bridge

import (
	"testing"
)

func TestBridge(t *testing.T) {

	d1 := Display{&StringDisplayImpl{"AAA"}}
	expect := "+---+\n|AAA|\n+---+\n"
	if result := d1.display(); result != expect {
		t.Errorf("Expect result to equal %s, but %s.\n", expect, result)
	}

	d2 := CountDisplay{&Display{&StringDisplayImpl{"BBB"}}}
	expect = "+---+\n|BBB|\n+---+\n"
	if result := d2.display(); result != expect {
		t.Errorf("Expect result to equal %s, but %s.\n", expect, result)
	}

	expect = "+---+\n|BBB|\n|BBB|\n+---+\n"
	if result := d2.multiDisplay(2); result != expect {
		t.Errorf("Expect result to equal %s, but %s.\n", expect, result)
	}
}
