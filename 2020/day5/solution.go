package day5

import (
	"fmt"
	"sort"
	"strings"

	"github.com/l33m4n123/adventOfCodeGo/utils"
)

type seat struct {
	ID     int
	row    int
	column int
}

// Solve runs the puzzle
func Solve(lines []string) {
	seats := parseSeats(lines)

	// Sort the seats by seatID
	sort.Slice(seats, func(i, j int) bool {
		return seats[i].ID < seats[j].ID
	})

	partOne(seats)
	partTwo(seats)
}

func partOne(seats []seat) {
	highestSeatID := seats[len(seats)-1].ID

	utils.PostSolution(5, 1, highestSeatID)
}

func partTwo(seats []seat) {
	expectedSeatID := seats[0].ID
	for _, seat := range seats {
		if seat.ID != expectedSeatID {
			utils.PostSolution(5, 2, expectedSeatID)
			return
		}
		expectedSeatID++
	}

	fmt.Println("Your seat was not found. Woopsie")
}

func parseSeats(lines []string) []seat {
	parsedSeats := []seat{}

	for _, line := range lines {
		lineSplice := strings.Split(line, "")
		seat := seat{}

		seat.row = parseInstructions(lineSplice[:7], 0, 127)
		seat.column = parseInstructions(lineSplice[7:], 0, 7)
		seat.ID = seat.row*8 + seat.column

		parsedSeats = append(parsedSeats, seat)
	}

	return parsedSeats
}

func parseInstructions(instructions []string, min int, max int) int {
	count := max + 1
	for _, instruction := range instructions {
		count /= 2
		switch instruction {
		case "F":
			fallthrough
		case "L":
			max -= count
			break
		case "B":
			fallthrough
		case "R":
			min += count
			break
		default:
			panic(fmt.Sprintf("%v is not a valid instruction", instruction))
		}
	}

	return max
}
