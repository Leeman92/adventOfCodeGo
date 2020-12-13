package day13

import (
	"fmt"
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
	var timeStamp int64 = 0
	// the multiplikation for all possible car times.. so we can now that if bus 2 can drive and we keep that score bus 1 can still go aswell thanks to modulo
	n1 := availableBusLinesPartTwo[0]

	for i := range availableBusLinesPartTwo {
		if i == 0 {
			continue
		}

		var j int64
		canDepart := false
		for !canDepart {
			j++
			possibleTimeStamp := timeStamp + int64(j)*n1                              // now we need to make sure to put n1 into it to make sure we take a time the previous bus(ses) can still depart at aswell
			canDepart = (possibleTimeStamp+int64(i))%availableBusLinesPartTwo[i] == 0 // if it is at 0 the bus can depart at the current timestamp
			if !canDepart {
				continue
			}
			timeStamp = possibleTimeStamp     // All the bus we checked so far can depart at the given time
			n1 *= availableBusLinesPartTwo[i] // As we now need to check when all the previous busses including this one can depart we increase the sum
		}
	}

	return timeStamp
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
