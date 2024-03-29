package day2

import (
	"github.com/l33m4n123/adventOfCodeGo/2021/submarine"
	"github.com/l33m4n123/adventOfCodeGo/2021/utils"
)

// Solve runs the puzzle
func Solve(lines []string) {
	sub := submarine.Submarine{}
	sub.PrepareNavigationComputer(lines, 2)
	sub.Steer()

	utils.PostSolution(2, 1, sub.GetPosition(false))
	utils.PostSolution(2, 2, sub.GetPosition(true))

	sub.DrawDepth()
}
