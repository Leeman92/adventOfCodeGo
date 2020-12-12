package day12

import (
	"fmt"
	"github.com/l33m4n123/adventOfCodeGo/2020/utils"
)

const (
	EAST = iota
	SOUTH
	WEST
	NORTH
)

type Boat struct {
	Location  utils.Coordinates
	Direction int
	Waypoint  utils.Coordinates // These are always relative to the boat
}

func (b *Boat) turn(direction string, degrees int) {
	if degrees%90 != 0 {
		panic(fmt.Sprintf("You want me to turn in in a weird angle (%v)\n", degrees))
	}

	amountOfTurns := (degrees / 90) % 4 // If the thing tells us to rotate by 3600 degrees. This would be 40 quarter turns or 10 full turns. which will end up at the same direction so we don't need to move

	newDirection := (b.Direction + amountOfTurns) % 4
	if direction == "L" {
		newDirection = (b.Direction - amountOfTurns) % 4
		if newDirection < 0 {
			newDirection += 4
		}
	}

	b.Direction = newDirection
}

func (b *Boat) move(direction string, amount int) {
	if direction == "F" {
		switch b.Direction {
		case EAST:
			direction = "E"
		case SOUTH:
			direction = "S"
		case WEST:
			direction = "W"
		case NORTH:
			direction = "N"
		}
	}

	switch direction {
	case "N":
		b.Location.Y -= amount
	case "S":
		b.Location.Y += amount
	case "E":
		b.Location.X += amount
	case "W":
		b.Location.X -= amount
	default:
		panic(fmt.Sprintf("It wants us to move in an unknown direction %v\n", direction))
	}
}

func (b *Boat) moveTowardsWaypoint(amount int) {
	// relative to me. South right means he is +x +y away
	// south east means +x -y and so on
	travelDistanceVertical := b.Waypoint.Y * amount
	travelDistanceHorizontal := b.Waypoint.X * amount

	if travelDistanceVertical < 0 {
		b.move("N", travelDistanceVertical*-1)
	} else {
		b.move("S", travelDistanceVertical)
	}

	if travelDistanceHorizontal < 0 {
		b.move("W", travelDistanceHorizontal*-1)
	} else {
		b.move("E", travelDistanceHorizontal)
	}
}

func (b *Boat) rotateWaypoint(direction string, degrees int) {
	if degrees%90 != 0 {
		panic(fmt.Sprintf("You want me to turn in in a weird angle (%v)\n", degrees))
	}

	amountOfTurns := (degrees / 90) % 4 // If the thing tells us to rotate by 3600 degrees. This would be 40 quarter turns or 10 full turns. which will end up at the same direction so we don't need to move

	if direction == "L" {
		amountOfTurns *= -1
	}

	switch amountOfTurns {
	case -3:
		fallthrough
	case 1:
		b.Waypoint.Y, b.Waypoint.X = b.Waypoint.X, b.Waypoint.Y*-1
	case -2:
		fallthrough
	case 2:
		b.Waypoint.X *= -1
		b.Waypoint.Y *= -1
	case -1:
		fallthrough
	case 3:
		b.Waypoint.Y, b.Waypoint.X = b.Waypoint.X*-1, b.Waypoint.Y
	}
}

func (b *Boat) moveWaypoint(direction string, amount int) {
	switch direction {
	case "N":
		b.Waypoint.Y -= amount
	case "S":
		b.Waypoint.Y += amount
	case "E":
		b.Waypoint.X += amount
	case "W":
		b.Waypoint.X -= amount
	default:
		panic(fmt.Sprintf("It wants us to move in an unknown direction %v\n", direction))
	}
}
