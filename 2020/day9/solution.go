package day9

import (
	"fmt"
	"sort"
	"time"

	"github.com/l33m4n123/adventOfCodeGo/2020/utils"
)

// Solve runs the puzzle
func Solve(input []string, test bool) {
	preamble := 25
	if test {
		preamble = 5
	}
	start := time.Now()
	intInput := utils.ConvertLinesToInt(input)
	solutionOne := partOne(intInput, preamble)

	solutionTwo := partTwo(intInput, preamble)

	utils.PostSolution(9, 1, solutionOne)
	utils.PostSolution(9, 2, solutionTwo)
	elapsed := time.Since(start)
	fmt.Printf("Solve took %s\n", elapsed)
}

func partOne(input []int, preamble int) int {
	for i := preamble; i < len(input); i++ {
		isValid := false
		for k := i - preamble; k < i; k++ {
			for j := k + 1; j < i; j++ {
				if input[k]+input[j] != input[i] {
					continue
				}
				isValid = true
			}

		}

		if !isValid {
			return input[i]
		}
	}

	return 0
}

func partTwo(input []int, preamble int) int {
	invalidIndex := getInvalidIndex(input, preamble)
	subInput := input[:invalidIndex]
	validRange := findValidRange(subInput, input[invalidIndex])

	sort.Slice(validRange, func(i, j int) bool {
		return validRange[i] < validRange[j]
	})

	return validRange[0] + validRange[len(validRange)-1]
}

func findValidRange(input []int, needle int) []int {
	for start := 0; start < len(input)-2; start++ {
		sum := input[start]
		for i := start + 1; i < len(input)-1; i++ {
			sum += input[i]
			if sum == needle {
				return input[start : i+1]
			}
		}
	}
	return input
}

func getInvalidIndex(input []int, preamble int) int {
	for i := preamble; i < len(input); i++ {
		isValid := false
		for k := i - preamble; k < i-1; k++ {
			for j := k + 1; j < i; j++ {
				if input[k]+input[j] != input[i] {
					continue
				}
				isValid = true
			}

		}

		if !isValid {
			return i
		}
	}

	return -1
}
