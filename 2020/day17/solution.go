package day17

import (
	"fmt"
	"strings"
	"time"

	"github.com/l33m4n123/adventOfCodeGo/2020/utils"
)

const inactiveCell = "."
const activeCell = "#"

// Cell is a data structure to keep each seats and floors
type Cell struct {
	active    bool
	nextState bool
	position  coordinates4D
}

func (c *Cell) updateState() {
	c.active = c.nextState
	c.nextState = false
}

type coordinates4D struct {
	x, y, z, w int
}

// Solve runs the puzzle. Looks like some sort of Conway's game of life.
func Solve(input []string) {
	start := time.Now()
	cells := parseInput(input)
	activeCounter := solve(cells, 0)
	utils.PostSolution(17, 1, activeCounter)
	elapsed := time.Since(start)
	fmt.Printf("%s elapsed to calculate Part 1\n", elapsed)

	start = time.Now()
	cells = parseInput(input)

	activeCounter = solve(cells, 1)
	utils.PostSolution(17, 2, activeCounter)
	elapsed = time.Since(start)
	fmt.Printf("%s elapsed to calculate Part 2\n", elapsed)
}

func solve(cells []*Cell, fourthDimension int) int {
	for cycle := 0; cycle < 6; cycle++ {
		cells = setupNewArea(cells, fourthDimension)
		cells = checkNewState(cells, fourthDimension)
		cells = updateCellState(cells)
	}

	//printCells(cells)

	return activeCellCount(cells)
}

func setupNewArea(cells []*Cell, fourthDimension int) []*Cell {
	alreadyChecked := map[int]map[int]map[int]map[int]bool{}
	for _, cell := range cells {
		for x := cell.position.x - 1; x <= cell.position.x+1; x++ {
			for y := cell.position.y - 1; y <= cell.position.y+1; y++ {
				for z := cell.position.z - 1; z <= cell.position.z+1; z++ {
					for w := cell.position.w - fourthDimension; w <= cell.position.w+fourthDimension; w++ {
						if alreadyChecked[x] == nil {
							alreadyChecked[x] = map[int]map[int]map[int]bool{}
						}
						if alreadyChecked[x][y] == nil {
							alreadyChecked[x][y] = map[int]map[int]bool{}
						}

						if alreadyChecked[x][y][z] == nil {
							alreadyChecked[x][y][z] = map[int]bool{}
						}

						if alreadyChecked[x][y][z][w] {
							continue
						}

						if cell.position.x == x && cell.position.y == y && cell.position.z == z && cell.position.w == w {
							continue
						}

						if !isCellSpot(cells, x, y, z, w) {
							newCell := Cell{}
							newCell.active = false
							newCell.position.x = x
							newCell.position.y = y
							newCell.position.z = z
							newCell.position.w = w
							cells = append(cells, &newCell)
						}
						alreadyChecked[x][y][z][w] = true
					}
				}
			}
		}
	}

	return cells
}

func isCellSpot(cells []*Cell, x, y, z, w int) bool {
	result := false
	for _, cell := range cells {
		if cell.position.x != x {
			continue
		}

		if cell.position.y != y {
			continue
		}

		if cell.position.z != z {
			continue
		}

		if cell.position.w != w {
			continue
		}

		result = true
		break
	}

	return result
}

func checkNewState(cells []*Cell, fourthDimension int) []*Cell {
	positionState := map[int]map[int]map[int]map[int]string{}
	var activeNeighours int
	for _, cell := range cells {
		activeNeighours, positionState = checkNeighbourCells(cells, cell, positionState, fourthDimension)
		if cell.active && activeNeighours == 2 {
			cell.nextState = true
			continue
		}
		if cell.active && activeNeighours == 3 {
			cell.nextState = true
			continue
		}
		if !cell.active && activeNeighours == 3 {
			cell.nextState = true
			continue
		}

		cell.nextState = false
	}
	return cells
}

func checkNeighbourCells(cells []*Cell,
	cell *Cell,
	positionState map[int]map[int]map[int]map[int]string,
	fourthDimension int) (int, map[int]map[int]map[int]map[int]string) {
	activeNeighours := 0
	counter := 0
	for x := cell.position.x - 1; x <= cell.position.x+1; x++ {
		for y := cell.position.y - 1; y <= cell.position.y+1; y++ {
			for z := cell.position.z - 1; z <= cell.position.z+1; z++ {
				for w := cell.position.w - fourthDimension; w <= cell.position.w+fourthDimension; w++ {
					if cell.position.x == x && cell.position.y == y && cell.position.z == z && cell.position.w == w {
						continue
					}

					counter++
					if positionState[x][y][z][w] == activeCell {
						activeNeighours++
						continue
					}

					if positionState[x][y][z][w] == inactiveCell {
						continue
					}

					if positionState[x] == nil {
						positionState[x] = map[int]map[int]map[int]string{}
					}

					if positionState[x][y] == nil {
						positionState[x][y] = map[int]map[int]string{}
					}

					if positionState[x][y][z] == nil {
						positionState[x][y][z] = map[int]string{}
					}

					neighbourState := getCellState(cells, x, y, z, w)
					positionState[x][y][z][w] = neighbourState
					if neighbourState == activeCell {
						activeNeighours++
					}
				}
			}
		}
	}

	return activeNeighours, positionState
}

func getCellState(cells []*Cell, x, y, z, w int) string {
	for _, cell := range cells {
		if cell.position.x != x ||
			cell.position.y != y ||
			cell.position.z != z ||
			cell.position.w != w {
			continue
		}

		if cell.active {
			return activeCell
		}

		return inactiveCell
	}

	return inactiveCell
}

func updateCellState(cells []*Cell) []*Cell {
	for idx, cell := range cells {
		cell.updateState()
		cells[idx] = cell
	}

	return cells
}

func activeCellCount(cells []*Cell) (count int) {
	for _, cell := range cells {
		if !cell.active {
			continue
		}

		count++
	}

	return count
}

func partTwo(seats [][]*Cell) int {
	return 0
}

func parseInput(input []string) []*Cell {
	cells := []*Cell{}
	for xIdx, line := range input {
		lineSplit := strings.Split(line, "")
		for yIdx, val := range lineSplit {
			cell := Cell{}
			cell.position.x = xIdx
			cell.position.y = yIdx
			cell.position.z = 0
			if val == activeCell {
				cell.active = true
			} else {
				cell.active = false
			}

			cells = append(cells, &cell)
		}
	}
	return cells
}

func printCells(cells []*Cell) {
	for _, cell := range cells {
		fmt.Println(cell)
	}
}
