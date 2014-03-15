package strategy

import (
	"testing"
)

func TestStrategy(t *testing.T) {
	player1 := player{name: "A", strategy: &winningStrategy{seed: 10}}
	player2 := player{name: "B", strategy: &winningStrategy{seed: 20}}

	hand1 := player1.nextHand()
	hand2 := player2.nextHand()

	if hand1.isStrongerThan(hand2) {
		player1.win()
		player2.lose()
	} else if hand1.isWeakerThan(hand2) {
		player1.lose()
		player2.win()
	} else {
		player1.even()
		player2.even()
	}
}
