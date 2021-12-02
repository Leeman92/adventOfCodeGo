package day2

import (
	"github.com/l33m4n123/adventOfCodeGo/2020/utils"
	"github.com/l33m4n123/adventOfCodeGo/2021/submarine"
)

// Solve runs the puzzle
func Solve(lines []string) {
	sub := submarine.Submarine{}
	sub.PrepareNavigationComputer(lines)
	sub.Steer()

	utils.PostSolution(2, 1, sub.GetPosition(false))
	utils.PostSolution(2, 2, sub.GetPosition(true))
}
