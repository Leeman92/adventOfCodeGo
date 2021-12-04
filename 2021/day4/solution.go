package day4

import (
	"github.com/l33m4n123/adventOfCodeGo/2021/submarine"
	"github.com/l33m4n123/adventOfCodeGo/2021/utils"
)

func Solve(lines []string) {
	sub := submarine.Submarine{}
	sub.GenerateBingoGame(lines)
	game := sub.BingoGame
	game.Play()
	partOne := game.FirstWinningBoard.CalculateUnmarkedSum() * game.FirstWinningNumber
	partTwo := game.LastWinningBoard.CalculateUnmarkedSum() * game.LastWinningNumber
	utils.PostSolution(4, 1, partOne)
	utils.PostSolution(4, 2, partTwo)
}
