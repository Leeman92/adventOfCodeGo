package day12

import (
	"regexp"
	"strconv"

	"github.com/l33m4n123/adventOfCodeGo/2020/utils"
)

// Solve runs the puzzle
func Solve(input []string) {
	boat := moveBoatPartOne(input)
	manhattenDistance := getManhattenDistance(boat)
	utils.PostSolution(12, 1, manhattenDistance)

	newBoat := moveBoatParTwo(input)
	newManhattenDistance := getManhattenDistance(newBoat)
	utils.PostSolution(12, 2, newManhattenDistance)
}

func moveBoatPartOne(input []string) *Boat {
	boat := newBoat()
	pattern := regexp.MustCompile(`(?P<instruction>.{1})(?P<amount>\d*)`)
	for _, line := range input {
		result := utils.FindStringSubmatchWithNamedMatches(pattern, line)
		amount, convErr := strconv.Atoi(result["amount"])
		if convErr != nil {
			panic(convErr)
		}
		if result["instruction"] == "L" || result["instruction"] == "R" {
			boat.turn(result["instruction"], amount)
		} else {
			boat.move(result["instruction"], amount)
		}
	}

	return &boat
}

func moveBoatParTwo(input []string) *Boat {
	boat := newBoat()
	pattern := regexp.MustCompile(`(?P<instruction>.{1})(?P<amount>\d*)`)
	for _, line := range input {
		result := utils.FindStringSubmatchWithNamedMatches(pattern, line)
		amount, convErr := strconv.Atoi(result["amount"])
		if convErr != nil {
			panic(convErr)
		}

		if result["instruction"] == "L" || result["instruction"] == "R" {
			boat.rotateWaypoint(result["instruction"], amount)
		} else if result["instruction"] == "F" {
			boat.moveTowardsWaypoint(amount)
		} else {
			boat.moveWaypoint(result["instruction"], amount)
		}
	}

	return &boat
}

func newBoat() Boat {
	b := Boat{}
	b.Location.X = 0
	b.Location.Y = 0
	b.Direction = EAST
	b.Waypoint.horizontal = 10
	b.Waypoint.vertical = -1

	return b
}

func getManhattenDistance(b *Boat) int {
	distanceTraversedHorizontal := b.Location.X
	distanceTraversedVertical := b.Location.Y

	if distanceTraversedHorizontal < 0 {
		distanceTraversedHorizontal *= -1
	}

	if distanceTraversedVertical < 0 {
		distanceTraversedVertical *= -1
	}

	return distanceTraversedHorizontal + distanceTraversedVertical
}
