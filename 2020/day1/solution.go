package day1

import (
	"fmt"

	"github.com/l33m4n123/adventOfCodeGo/2020/utils"
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
				fmt.Printf("%d + %d = 2020\n", x, y)
				fmt.Printf("Your answer for Part 1 is %d\n", x*y)
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
					fmt.Printf("%d + %d + %d = 2020\n", x, y, z)
					fmt.Printf("Your answer for Part 2 is %d\n", x*y*z)
					return
				}
			}
		}
	}
}
