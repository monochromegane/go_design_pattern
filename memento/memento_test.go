package memento

import (
	"testing"
)

func TestMemento(t *testing.T) {

	game := &Game{100}
	memento := game.CreateMemento()

	game.Money = 0

	game.RestoreMemento(memento)

	if game.Money != 100 {
		t.Errorf("Expect money to equal %s, but %s\n", 100, game.Money)
	}

}
