package day13

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/l33m4n123/adventOfCodeGo/2020/utils"
)

var earliestDeparture int
var availableBusLines []int
var availableBusLinesPartTwo map[int]int64
var testRun bool

func Solve(input []string, test bool) {
	testRun = test
	parseInput(input)
	solutionPartOne := partOne()
	utils.PostSolution(13, 1, solutionPartOne)

	start := time.Now()
	timeStampPartTwo := partTwo()
	utils.PostSolution(13, 2, timeStampPartTwo)
	elapsed := time.Since(start)
	fmt.Printf("It took %s to calculate part 2\n", elapsed)
}

func partOne() int {
	earliestDepartureTimestamp, possibleBus := getEarliestDepartureTimestamp()
	return possibleBus * (earliestDepartureTimestamp - earliestDeparture)
}

func partTwo() int64 {
	return getSolutionPartTwo()
}

func getSolutionPartTwo() int64 {

	// Suppose bus B arrives at index I in the list
	// B should depart at time T+I.
	// T+I % B == 0
	// that means
	// T % B == -I
	// because I want positive values
	// T % B == (B -(I%B))%B

	// For the chinese remainder theorem we need an N that is the product of the starting times. Because 0 < T < N according to the crt
	product := int64(1)
	remainders := map[int64]int64{}
	for index, bus := range availableBusLinesPartTwo {
		remainders[bus-(int64(index)%bus)%bus] = bus
		product *= int64(bus)
	}

	answer := int64(0)
	for I, B := range remainders {
		newProductFloat := float64(product) / math.Max(float64(B), 1)
		newProduct := int64(newProductFloat)
		// newProduct is the product of all other bus starting times
		// if we add a multiple of NI to T, it won't affect when the other busses arive modulo T,
		// since newProduct is a multiple of each other busses

		// We need to find a multiple of newProduct so that (x*newProduct)%B == I
		// first find modInverse so that (modInverse*newProduct)%B == 1
		// then (index * modInverse * newProduct)%B == 1

		// First check if we only have primes
		if utils.GetGCD(newProduct, B) != 1 {
			panic("We have no primes. CRT only works with primes")
		}
		modInverse := utils.ModInverse(newProduct, B)
		time := I * modInverse * newProduct
		answer += time
	}

	answer %= product

	return answer
}

func getEarliestDepartureTimestamp() (int, int) {
	for i := earliestDeparture; i < earliestDeparture+20; i++ {
		for _, busLine := range availableBusLines {
			if busLine == -1 {
				continue
			}
			if i%busLine == 0 {
				return i, busLine
			}
		}
	}

	return -1, -1
}

func parseInput(input []string) {
	if len(input) != 2 {
		panic("The input is malformed")
	}

	intVal, convErr := strconv.Atoi(input[0])
	if convErr != nil {
		panic(convErr)
	}

	earliestDeparture = intVal

	busLines := strings.Split(input[1], ",")
	availableBusLinesPartTwo = make(map[int]int64, len(busLines))
	for index, busLine := range busLines {
		if busLine == "x" {
			continue
		}

		busIntVal, busConvErr := strconv.Atoi(busLine)
		if busConvErr != nil {
			panic(busConvErr)
		}

		availableBusLines = append(availableBusLines, busIntVal)
		availableBusLinesPartTwo[index] = int64(busIntVal)
	}
}
