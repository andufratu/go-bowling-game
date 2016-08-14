package go_bowling_game

type BowlingGame struct {
	CurrentRoll int
	Rolls [21]int
}

func (game *BowlingGame) Roll(pins int) {
	game.Rolls[game.CurrentRoll] = pins
	game.CurrentRoll += 1
}

func (game *BowlingGame) Score() (score int) {
	score = 0
	frameIndex := 0
	for i := 0; i < 10; i++ {
		if game.isStrike(frameIndex) {
			score += 10 + game.strikeBonus(frameIndex)
			frameIndex += 1
		} else if game.isSpare(frameIndex) {
			score += 10 + game.spareBonus(frameIndex)
			frameIndex += 2
		} else {
			score += game.sumOfBallsInFrame(frameIndex)
			frameIndex += 2
		}
	}
	return
}

func (game *BowlingGame) isSpare(frameIndex int) bool {
	return game.Rolls[frameIndex] + game.Rolls[frameIndex + 1] == 10
}

func (game *BowlingGame) isStrike(frameIndex int) bool {
	return game.Rolls[frameIndex] == 10
}

func (game *BowlingGame) spareBonus(frameIndex int) int {
	return game.Rolls[frameIndex + 2]
}

func (game *BowlingGame) strikeBonus(frameIndex int) int {
	return game.Rolls[frameIndex + 1] + game.Rolls[frameIndex + 2]
}

func (game *BowlingGame) sumOfBallsInFrame(frameIndex int) int {
	return game.Rolls[frameIndex] + game.Rolls[frameIndex + 1]
}
