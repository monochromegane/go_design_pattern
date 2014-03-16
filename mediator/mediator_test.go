package mediator

import (
	"testing"
)

func TestMediator(t *testing.T) {
	loginForm := newLoginForm()

	state := loginForm.button.enabled
	if state {
		t.Errorf("Expect state to false, but true.\n")
	}

	loginForm.radioButton.check(true)

	state = loginForm.button.enabled
	if !state {
		t.Errorf("Expect state to true, but false.\n")
	}

}
