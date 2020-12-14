package day10

import (
	"sort"
	"time"

	"github.com/l33m4n123/adventOfCodeGo/utils"
)

func Solve(input []string) {
	adapters := utils.ConvertLinesToInt(input)
	adapters = append(adapters, 0) // Outlet is 0
	sort.Slice(adapters, func(i, j int) bool {
		return adapters[i] < adapters[j]
	})
	adapters = append(adapters, adapters[len(adapters)-1]+3)
	solutionOne := partOne(adapters)
	solutionTwo := partTwo(adapters)

	utils.PostSolution(10, 1, solutionOne)
	start := time.Now()
	utils.PostSolution(10, 2, solutionTwo)
	elapsed := time.Since(start)

	utils.PostSolution(10, 2, elapsed)
}

func partOne(input []int) int {
	oneJoltDifference := 0
	threeJoltDifference := 0
	for i := 1; i < len(input); i++ {
		difference := input[i] - input[i-1]
		switch difference {
		case 1:
			oneJoltDifference++
		case 3:
			threeJoltDifference++
		}
	}

	return oneJoltDifference * threeJoltDifference
}

func partTwo(input []int) int {
	cache := map[int]int{}
	possibleCombinations := findIterationCount(0, input, cache)
	return possibleCombinations
}

func findIterationCount(currIndex int, input []int, cache map[int]int) int {
	if currIndex == len(input)-1 {
		return 1
	}

	if val, ok := cache[currIndex]; ok {
		return val
	}
	count := 0
	for i := currIndex + 1; i < len(input); i++ {
		if input[i]-input[currIndex] <= 3 {
			count += findIterationCount(i, input, cache)
		}
	}
	cache[currIndex] = count
	return count
}
