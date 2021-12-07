package day7

import (
	"github.com/l33m4n123/adventOfCodeGo/2021/utils"
	"math"
	"sort"
)

func Solve(lines []string) {
	nodes := utils.ParseInputToIntSlice(lines)
	sort.Ints(nodes)

	part1 := solvePart1(nodes)
	utils.PostSolution(7, 1, part1)
	part2 := solvePart2(nodes)

	utils.PostSolution(7, 1, part2)
}

func solvePart1(nodes []int) (smallestDistance int) {
	minVal := nodes[0]
	maxVal := nodes[len(nodes)-1]
	smallestDistance = math.MaxInt
	for i := minVal; i <= maxVal; i++ {
		currentDistance := 0
		for _, node := range nodes {
			distance := node - i
			if distance < 0 {
				distance *= -1
			}
			currentDistance += distance
		}
		if currentDistance < smallestDistance {
			smallestDistance = currentDistance
		}
	}

	return smallestDistance
}

func solvePart2(nodes []int) (smallestDistance int) {
	minVal := nodes[0]
	maxVal := nodes[len(nodes)-1]
	sumFormula := make(map[int]int)
	smallestDistance = math.MaxInt
	for i := minVal; i <= maxVal; i++ {
		currentDistance := 0
		for _, node := range nodes {
			distance := node - i
			if distance < 0 {
				distance *= -1
			}

			currentDistance += calcSumFormula(distance, 0, sumFormula)
		}
		if currentDistance < smallestDistance {
			smallestDistance = currentDistance
		}
	}

	return smallestDistance
}

func calcSumFormula(startValue int, sum int, cache map[int]int) int {
	if startValue <= 1 {
		return startValue
	}

	val, ok := cache[startValue]
	if ok {
		return val
	}

	cache[startValue] = startValue + calcSumFormula(startValue-1, sum, cache)
	return cache[startValue]
}
