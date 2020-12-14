package day1

import (
	"github.com/l33m4n123/adventOfCodeGo/utils"
)

// Solve runs the puzzle
func Solve(lines []string) {
	intLines := utils.ConvertLinesToInt(lines)

	partOne(intLines)
	partTwo(intLines)
}

func partOne(lines []int) {
	for outerKey, x := range lines {
		for innerKey, y := range lines {
			if innerKey == outerKey || innerKey <= outerKey {
				continue
			}

			if x+y == 2020 {
				utils.PostSolution(1, 1, x*y)
				return
			}
		}
	}
}

func partTwo(lines []int) {
	for xKey, x := range lines {
		for yKey, y := range lines {
			for zKey, z := range lines {
				if xKey == yKey ||
					xKey == zKey ||
					yKey == zKey ||
					yKey < xKey ||
					zKey < yKey {
					continue
				}

				if x+y+z == 2020 {
					utils.PostSolution(1, 2, x*y*z)
					return
				}
			}
		}
	}
}
