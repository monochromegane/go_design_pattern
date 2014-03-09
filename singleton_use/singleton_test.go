package use_singleton

import (
	"github.com/monochromegane/go_design_pattern/singleton"
	"testing"
)

func TestSingleton(t *testing.T) {

	// cannot refer to unexported name singleton.singleton
	// instance := singleton.singleton{}

	instance1 := singleton.GetInstance()
	instance2 := singleton.GetInstance()

	if instance1 != instance2 {
		t.Errorf("Expect instance to equal, but not equal.\n")
	}

}
