package template_method

import (
	"testing"
)

func TestCharDisplay(t *testing.T) {
	display := &CharDisplay{char: 'A'}
	result := display.display(display)
	if result != "<<AAAAA>>" {
		t.Errorf("Expect result to %s, but %s.\n", "<<AAAAA>>", result)
	}
}

func TestStringDisplay(t *testing.T) {
	expect := "+-------+\n| ABCDE |\n| ABCDE |\n| ABCDE |\n| ABCDE |\n| ABCDE |\n+-------+\n"
	display := &StringDisplay{str: "ABCDE"}
	result := display.display(display)
	if result != expect {
		t.Errorf("Expect result to \n%s, but \n%s.\n", expect, result)
	}
}
