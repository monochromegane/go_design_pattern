package state

import (
	"testing"
)

func TestState(t *testing.T) {

	safeFrame := &SafeFrame{State: GetDayInstance()}

	hours := []int{8, 9, 16, 17}

	for _, hour := range hours {
		safeFrame.SetClock(hour)
		safeFrame.Use()
	}

	expect := "night day day night"
	if safeFrame.GetLog() != expect {
		t.Errorf("Expect log to equal %s, but %s\n", expect, safeFrame.GetLog())
	}

}
