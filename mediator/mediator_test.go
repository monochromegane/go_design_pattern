package mediator

import (
	"testing"
)

func TestMediator(t *testing.T) {
	loginForm := NewLoginForm()

	state := loginForm.Button.Enabled
	if state {
		t.Errorf("Expect state to false, but true.\n")
	}

	loginForm.RadioButton.Check(true)

	state = loginForm.Button.Enabled
	if !state {
		t.Errorf("Expect state to true, but false.\n")
	}

}
