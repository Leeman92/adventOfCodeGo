package day15

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/l33m4n123/adventOfCodeGo/2020/utils"
)

func Solve(input []string) {
	result := solve(input[0], 2020)
	utils.PostSolution(15, 1, result)

	start := time.Now()
	result = solve(input[0], 30000000)
	utils.PostSolution(15, 2, result)
	elapsed := time.Since(start)

	fmt.Printf("It took %s to process part 2\n", elapsed)
}

func solve(line string, number int) int {
	memory := make(map[int]int)
	lastSaid := 0
	lineSplit := strings.Split(line, ",")

	for key, said := range lineSplit {
		saidAsInt, err := strconv.Atoi(said)
		if err != nil {
			panic(err)
		}
		lastSaid = saidAsInt
		if key < len(lineSplit)-1 {
			memory[saidAsInt] = key + 1
		}
	}

	for i := len(lineSplit) + 1; i <= number; i++ {
		sayingNow := memory[lastSaid]
		if sayingNow > 0 {
			sayingNow = i - 1 - sayingNow
		}
		memory[lastSaid] = i - 1
		lastSaid = sayingNow
	}

	return lastSaid
}
