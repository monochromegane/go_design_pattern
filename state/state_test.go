package state

import (
	"testing"
)

func TestState(t *testing.T) {

	safeFrame := &safeFrame{state: GetDayInstance()}

	hours := []int{8, 9, 16, 17}

	for _, hour := range hours {
		safeFrame.setClock(hour)
		safeFrame.use()
	}

	expect := "night day day night"
	if safeFrame.getLog() != expect {
		t.Errorf("Expect log to equal %s, but %s\n", expect, safeFrame.getLog())
	}

}
