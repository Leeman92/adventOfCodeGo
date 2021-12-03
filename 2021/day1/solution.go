package day1

import (
	"fmt"
	"github.com/l33m4n123/adventOfCodeGo/2021/utils"
	"time"
)

// Solve runs the puzzle
func Solve(lines []string) {
	intLines := utils.ConvertLinesToInt(lines)

	partOne(intLines)
	partTwo(intLines)
}

func partOne(lines []int) {
	lastVal := 0
	counter := 0
	for index, val := range lines {
		if index == 0 {
			lastVal = val
			continue
		}
		if lastVal < val {
			counter++
		}
		lastVal = val
	}
	utils.PostSolution(1, 1, counter)
}

func partTwo(lines []int) {
	createPairs(lines)
}

func createPairs(lines []int) {
	counter := 0
	start := time.Now()
	lastSum := 0
	for i := 0; i < len(lines); i++ {
		if i+3 >= len(lines) {
			break
		}
		sumOne := 0
		intermediateSum := 0
		if lastSum > 0 {
			sumOne = lastSum
			intermediateSum = sumOne - lines[i]
		} else {
			intermediateSum = lines[i+1] + lines[i+2]
			sumOne = lines[i] + intermediateSum
		}

		sumTwo := intermediateSum + lines[i+3]
		lastSum = sumTwo
		if sumTwo > sumOne {
			counter++
		}
	}
	elapsed := time.Since(start)
	fmt.Printf("%s elapsed to calculate Part 2\n", elapsed)

	utils.PostSolution(1, 2, counter)
}
