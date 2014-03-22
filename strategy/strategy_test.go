package strategy

import (
	"testing"
)

func TestStrategy(t *testing.T) {
	player1 := Player{Name: "A", Strategy: &winningStrategy{seed: 10}}
	player2 := Player{Name: "B", Strategy: &winningStrategy{seed: 20}}

	hand1 := player1.NextHand()
	hand2 := player2.NextHand()

	if hand1.IsStrongerThan(hand2) {
		player1.Win()
		player2.Lose()
	} else if hand1.IsWeakerThan(hand2) {
		player1.Lose()
		player2.Win()
	} else {
		player1.Even()
		player2.Even()
	}
}
