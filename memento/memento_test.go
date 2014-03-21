package memento

import (
	"testing"
)

func TestMemento(t *testing.T) {

	game := &game{100}
	memento := game.createMemento()

	game.money = 0

	game.restoreMemento(memento)

	if game.money != 100 {
		t.Errorf("Expect money to equal %s, but %s\n", 100, game.money)
	}

}
