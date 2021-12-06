package day6

import (
	"github.com/l33m4n123/adventOfCodeGo/2021/utils"
	"strconv"
	"strings"
)

func Solve(lines []string) {
	lineSlice := strings.Split(lines[0], ",")
	fishes := prepareFishes(lineSlice)

	fishes = simulateDays(80, fishes)
	utils.PostSolution(6, 1, count(fishes))

	fishes = simulateDays(256-80, fishes)
	utils.PostSolution(6, 2, count(fishes))
}

func count(fishes map[int]int) (counter int) {
	for i := 0; i < 9; i++ {
		counter += fishes[i]
	}

	return counter
}

func prepareFishes(line []string) map[int]int {

	fishes := make(map[int]int, 9)
	for _, val := range line {
		number, err := strconv.Atoi(val)
		if err != nil {
			panic(err)
		}
		fishes[number] += 1
	}

	for i := 0; i < 9; i++ {
		_, ok := fishes[i]
		if !ok {
			fishes[i] = 0
		}
	}

	return fishes
}

func simulateDays(days int, fishes map[int]int) map[int]int {
	for i := 0; i < days; i++ {
		fishesGivingBirth := 0
		for day := 0; day < 8; day++ {
			if day == 0 {
				fishesGivingBirth = fishes[day]
			}
			fishes[day] = fishes[day+1]
		}
		fishes[6] += fishesGivingBirth
		fishes[8] = fishesGivingBirth
	}

	return fishes
}
