package facade

import (
	"testing"
)

func TestFacade(t *testing.T) {
	maker := pageMaker{}
	result := maker.makeWelcomePage("a@a.com")
	expect := "# Welcome to a's page!"

	if result != expect {
		t.Errorf("Expect result to equal %s, but %s.", expect, result)
	}
}
