package day11

import (
	"fmt"
	"strings"
	"time"

	"github.com/l33m4n123/adventOfCodeGo/2020/utils"
)

const floor = "."
const emtpySeat = "L"
const occupiedSeat = "#"

// Position is a data structure to keep each seats and floors
type Position struct {
	directNeighbours []*Position
	lineNeighbours   []*Position
	occupied         bool
	isSeat           bool
	nextOccupation   bool
	position         coordinates
}

// UpdateOccupation updates the positions occupation based on the rules
func (pos *Position) UpdateOccupation(dry bool) bool {
	if !dry {
		originalOccupation := pos.occupied
		pos.occupied = pos.nextOccupation
		return originalOccupation != pos.nextOccupation
	}

	if !pos.isSeat {
		pos.nextOccupation = false
		return false
	}

	occupiedNeighbours := 0
	for _, neighbours := range pos.directNeighbours {
		if neighbours.isSeat && neighbours.occupied {
			occupiedNeighbours++
		}
	}

	if !pos.occupied && occupiedNeighbours == 0 {
		pos.nextOccupation = true
		return true
	}

	if pos.occupied && occupiedNeighbours >= 4 {
		pos.nextOccupation = false
		return false
	}

	return false
}

// UpdateOccupationPartTwo updates the positions occupation based on the new rules
func (pos *Position) UpdateOccupationPartTwo(dry bool) bool {
	if !dry {
		originalOccupation := pos.occupied
		pos.occupied = pos.nextOccupation
		return originalOccupation != pos.nextOccupation
	}

	if !pos.isSeat {
		pos.nextOccupation = false
		return false
	}

	occupiedNeighbours := 0
	for _, neighbours := range pos.lineNeighbours {
		if neighbours.isSeat && neighbours.occupied {
			occupiedNeighbours++
		}
	}

	if !pos.occupied && occupiedNeighbours == 0 {
		pos.nextOccupation = true
		return true
	}

	if pos.occupied && occupiedNeighbours >= 5 {
		pos.nextOccupation = false
		return false
	}

	return false
}

// AddDirectNeighbour adds a new neighbour to the seat.
func (pos *Position) AddDirectNeighbour(neighbour *Position) {
	pos.directNeighbours = append(pos.directNeighbours, neighbour)
}

// AddLineNeighbour adds a new neighbour to the seat
func (pos *Position) AddLineNeighbour(neighbour *Position) {
	pos.lineNeighbours = append(pos.lineNeighbours, neighbour)
}

type coordinates struct {
	row, col int
}

// Solve runs the puzzle. Looks like some sort of Conway's game of life.
func Solve(input []string) {
	start := time.Now()
	seats := parseInput(input)

	occupationCounter := partOne(seats)
	seats = resetOccupation(seats)
	utils.PostSolution(11, 1, occupationCounter)

	occupationCounter = partTwo(seats)
	elapsed := time.Since(start)
	utils.PostSolution(11, 2, occupationCounter)

	fmt.Printf("%s elapsed to update seats\n", elapsed)
}

func partOne(seats [][]*Position) int {
	change := true
	for change {
		change = false
		// we first need to run dry over it to mark the possible changes and after that, do the actual change
		for _, row := range seats {
			for _, seat := range row {
				seat.UpdateOccupation(true)
			}
		}

		for _, row := range seats {
			for _, seat := range row {
				if seat.UpdateOccupation(false) {
					change = true
				}
			}
		}
	}

	occupationCounter := 0
	for _, row := range seats {
		for _, seat := range row {
			if seat.occupied {
				occupationCounter++
			}
		}
	}

	return occupationCounter
}

func partTwo(seats [][]*Position) int {
	change := true
	for change {

		change = false
		// we first need to run dry over it to mark the possible changes and after that, do the actual change
		for _, row := range seats {
			for _, seat := range row {
				seat.UpdateOccupationPartTwo(true)
			}
		}

		for _, row := range seats {
			for _, seat := range row {
				if seat.UpdateOccupationPartTwo(false) {
					change = true
				}
			}
		}
	}

	occupationCounter := 0
	for _, row := range seats {
		for _, seat := range row {
			if seat.occupied {
				occupationCounter++
			}
		}
	}

	return occupationCounter
}

func resetOccupation(seats [][]*Position) [][]*Position {
	for _, row := range seats {
		for _, seat := range row {
			seat.occupied = false
		}
	}

	return seats
}

func parseInput(input []string) [][]*Position {
	var rowsAndCols [][]string
	rowCount := len(input)
	colCount := len(input[0])

	for _, row := range input {
		rowSlice := strings.Split(row, "")
		rowsAndCols = append(rowsAndCols, rowSlice)
	}

	positions := createPositions(rowsAndCols, rowCount, colCount)

	return positions
}

func createPositions(rowsAndCols [][]string, rowCount, colCount int) [][]*Position {
	positions := make([][]*Position, rowCount)

	for col := range positions {
		positions[col] = make([]*Position, colCount)
	}

	for row := 0; row < rowCount; row++ {
		for col := 0; col < colCount; col++ {
			pos := Position{}
			pos.position = coordinates{row, col}
			pos.occupied = false
			pos.isSeat = rowsAndCols[row][col] != floor
			positions[row][col] = &pos
		}
	}

	positions = populateDirectNeighbours(positions, rowCount, colCount)
	positions = populateLineNeighbours(positions, rowCount, colCount)

	return positions
}

func populateDirectNeighbours(positions [][]*Position, maxRow, maxCol int) [][]*Position {
	for rowCount, row := range positions {
		for colCount, pos := range row {
			if colCount > 0 {
				pos.AddDirectNeighbour(positions[rowCount][colCount-1])
				if rowCount > 0 {
					pos.AddDirectNeighbour(positions[rowCount-1][colCount-1])
				}
				if rowCount < maxRow-1 {
					pos.AddDirectNeighbour(positions[rowCount+1][colCount-1])
				}
			}

			if colCount < maxCol-1 {
				pos.AddDirectNeighbour(positions[rowCount][colCount+1])
				if rowCount > 0 {
					pos.AddDirectNeighbour(positions[rowCount-1][colCount+1])
				}
				if rowCount < maxRow-1 {
					pos.AddDirectNeighbour(positions[rowCount+1][colCount+1])
				}
			}

			if rowCount > 0 {
				pos.AddDirectNeighbour(positions[rowCount-1][colCount])
			}
			if rowCount < maxRow-1 {
				pos.AddDirectNeighbour(positions[rowCount+1][colCount])
			}

			positions[rowCount][colCount] = pos
		}
	}

	return positions
}

func populateLineNeighbours(positions [][]*Position, maxRow, maxCol int) [][]*Position {
	for rowCount, row := range positions {
		for colCount, pos := range row {
			inspectRowUp(pos, positions, 0)
			inspectRowDown(pos, positions, maxRow)
			inspectColLeft(pos, positions, 0)
			inspectColRight(pos, positions, maxCol)
			inspectDiagLeftUp(pos, positions, 0, 0)
			inspectDiagUpRight(pos, positions, 0, maxCol)
			inspectDiagRightDown(pos, positions, maxRow, maxCol)
			inspectDiagDownLeft(pos, positions, maxRow, 0)

			positions[rowCount][colCount] = pos
		}
	}

	return positions
}

func inspectRowUp(pos *Position, positions [][]*Position, minRow int) {
	for inspectedRow := pos.position.row - 1; inspectedRow >= minRow; inspectedRow-- {
		posSeen := positions[inspectedRow][pos.position.col]
		if !posSeen.isSeat {
			continue
		}

		pos.AddLineNeighbour(posSeen)
		break
	}
}

func inspectRowDown(pos *Position, positions [][]*Position, maxRow int) {
	for inspectedRow := pos.position.row + 1; inspectedRow < maxRow; inspectedRow++ {
		posSeen := positions[inspectedRow][pos.position.col]
		if !posSeen.isSeat {
			continue
		}

		pos.AddLineNeighbour(posSeen)
		break
	}
}

func inspectColLeft(pos *Position, positions [][]*Position, minCol int) {
	for inspectedCol := pos.position.col - 1; inspectedCol >= minCol; inspectedCol-- {
		posSeen := positions[pos.position.row][inspectedCol]
		if !posSeen.isSeat {
			continue
		}

		pos.AddLineNeighbour(posSeen)
		break
	}
}

func inspectColRight(pos *Position, positions [][]*Position, maxCol int) {
	for inspectedCol := pos.position.col + 1; inspectedCol < maxCol; inspectedCol++ {
		posSeen := positions[pos.position.row][inspectedCol]
		if !posSeen.isSeat {
			continue
		}

		pos.AddLineNeighbour(posSeen)
		break
	}
}

func inspectDiagLeftUp(pos *Position, positions [][]*Position, minRow, minCol int) {
	for inspectedRow := pos.position.row - 1; inspectedRow >= minRow; inspectedRow-- {
		difference := pos.position.row - inspectedRow
		newCol := pos.position.col - difference
		if newCol < minCol {
			break
		}

		posSeen := positions[inspectedRow][pos.position.col-difference]
		if !posSeen.isSeat {
			continue
		}

		pos.AddLineNeighbour(posSeen)
		break
	}
}

func inspectDiagUpRight(pos *Position, positions [][]*Position, minRow, maxCol int) {
	for inspectedRow := pos.position.row - 1; inspectedRow >= minRow; inspectedRow-- {
		difference := pos.position.row - inspectedRow
		newCol := pos.position.col + difference
		if newCol >= maxCol {
			break
		}

		posSeen := positions[inspectedRow][newCol]
		if !posSeen.isSeat {
			continue
		}

		pos.AddLineNeighbour(posSeen)
		break
	}
}

func inspectDiagRightDown(pos *Position, positions [][]*Position, maxRow, maxCol int) {
	for inspectedRow := pos.position.row + 1; inspectedRow < maxRow; inspectedRow++ {
		difference := inspectedRow - pos.position.row
		newCol := pos.position.col + difference
		if newCol >= maxCol {
			break
		}

		posSeen := positions[inspectedRow][newCol]
		if !posSeen.isSeat {
			continue
		}

		pos.AddLineNeighbour(posSeen)
		break
	}
}

func inspectDiagDownLeft(pos *Position, positions [][]*Position, maxRow, minCol int) {
	for inspectedRow := pos.position.row + 1; inspectedRow < maxRow; inspectedRow++ {
		difference := inspectedRow - pos.position.row
		newCol := pos.position.col - difference
		if newCol < minCol {
			break
		}
		posSeen := positions[inspectedRow][newCol]
		if !posSeen.isSeat {
			continue
		}

		pos.AddLineNeighbour(posSeen)
		break
	}
}
