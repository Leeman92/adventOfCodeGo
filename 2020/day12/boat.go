package day12

import "fmt"

const (
	EAST = iota
	SOUTH
	WEST
	NORTH
)

type Boat struct {
	Location  Coordinates
	Direction int
	Waypoint  RelativeCoordinates
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
	travelDistanceVertical := b.Waypoint.vertical * amount
	travelDistanceHorizontal := b.Waypoint.horizontal * amount

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

	horizontal := b.Waypoint.horizontal
	vertical := b.Waypoint.vertical
	switch amountOfTurns {
	case -3:
		fallthrough
	case 1:
		b.Waypoint.vertical, b.Waypoint.horizontal = horizontal, vertical*-1
	case -2:
		fallthrough
	case 2:
		b.Waypoint.horizontal *= -1
		b.Waypoint.vertical *= -1
	case -1:
		fallthrough
	case 3:
		b.Waypoint.vertical, b.Waypoint.horizontal = horizontal*-1, vertical
	}
}

func (b *Boat) moveWaypoint(direction string, amount int) {
	switch direction {
	case "N":
		b.Waypoint.vertical -= amount
	case "S":
		b.Waypoint.vertical += amount
	case "E":
		b.Waypoint.horizontal += amount
	case "W":
		b.Waypoint.horizontal -= amount
	default:
		panic(fmt.Sprintf("It wants us to move in an unknown direction %v\n", direction))
	}
}

type Coordinates struct {
	X, Y int
}

type RelativeCoordinates struct {
	horizontal, vertical int
}
