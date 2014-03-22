package template_method

import (
	"testing"
)

func TestCharDisplay(t *testing.T) {
	display := &CharDisplay{Char: 'A'}
	result := display.Display(display)
	if result != "<<AAAAA>>" {
		t.Errorf("Expect result to %s, but %s.\n", "<<AAAAA>>", result)
	}
}

func TestStringDisplay(t *testing.T) {
	expect := "+-------+\n| ABCDE |\n| ABCDE |\n| ABCDE |\n| ABCDE |\n| ABCDE |\n+-------+\n"
	display := &StringDisplay{Str: "ABCDE"}
	result := display.Display(display)
	if result != expect {
		t.Errorf("Expect result to \n%s, but \n%s.\n", expect, result)
	}
}
