package bowling_game

import (
	"testing"
	"fmt"
)

var game *BowlingGame

func TestGutterGame(t *testing.T) {
	initGame()
	rollMany(20, 0)

	score := game.Score()
	if score != 0 {
		fmt.Printf("Score was supposed to be 0 for gutter game, it's actually %v\n", score)
		t.Fail()
	}
}

func TestAllOnesGame(t *testing.T) {
	initGame()
	rollMany(20, 1)

	score := game.Score()
	if score != 20 {
		fmt.Printf("Score was supposed to be 20 for all ones, it's actually %v\n", score)
		t.Fail()
	}
}

func TestOneSpare(t *testing.T) {
	initGame()
	rollSpare()
	game.Roll(3)
	rollMany(17, 0)

	score := game.Score()
	if score != 16 {
		fmt.Printf("Score was supposed to be 16 for one spare, it's actually %v\n", score)
		t.Fail()
	}
}

func TestOneStrike(t *testing.T) {
	initGame()
	rollStrike()
	game.Roll(3)
	game.Roll(4)
	rollMany(16, 0)

	score := game.Score()
	if score != 24 {
		fmt.Printf("Score was supposed to be 24 for one strike, it's actually %v\n", score)
		t.Fail()
	}
}

func TestPerfectGame(t *testing.T) {
	initGame()
	rollMany(12, 10)

	score := game.Score()
	if score != 300 {
		fmt.Printf("Score was supposed to be 300 for perfect game, it's actually %v\n", score)
		t.Fail()
	}
}

func initGame() {
	game = new(BowlingGame)
}

func rollMany(count, pins int) {
	for i := 0; i < count; i++ {
		game.Roll(pins)
	}
}

func rollSpare() {
	game.Roll(5)
	game.Roll(5)
}

func rollStrike() {
	game.Roll(10)
}
