package day3

import (
	"github.com/l33m4n123/adventOfCodeGo/2021/submarine"
	"github.com/l33m4n123/adventOfCodeGo/2021/utils"
)

// Solve runs the puzzle
func Solve(lines []string) {
	sub := submarine.Submarine{}
	sub.RunDiagnostic(lines)

	utils.PostSolution(3, 1, sub.PowerDraw)
	utils.PostSolution(3, 2, sub.LifeSupportRating)
}
