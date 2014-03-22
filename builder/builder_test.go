package builder

import (
	"testing"
)

func TestBuilder(t *testing.T) {

	expect := "# Title\n## String\n- Item1\n- Item2\n\n"

	director := Director{&TextBuilder{}}
	result := director.Construct()

	if result != expect {
		t.Errorf("Expect result to %s, but %s", result, expect)
	}

}
