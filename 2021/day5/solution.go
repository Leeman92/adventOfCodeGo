package day5

import (
	"fmt"
	"github.com/l33m4n123/adventOfCodeGo/2021/utils"
	"strconv"
	"strings"
	"time"
)

type hydrothermalVents struct {
	VisitedLocation map[int]map[int]int
	maxX            int
	maxY            int
}

func (v *hydrothermalVents) Overlap() (counter int) {
	for y := 0; y <= v.maxY; y++ {
		for x := 0; x <= v.maxX; x++ {
			val, ok := v.VisitedLocation[x][y]
			if !ok {
				continue
			}
			if val >= 2 {
				counter++
			}
		}
	}

	return counter
}

type Line struct {
	StartX, StartY, EndX, EndY int
	diagonal                   bool
}

func Solve(input []string) {
	var lines []Line
	start := time.Now()
	for _, val := range input {
		newLine := parseLine(val)
		lines = append(lines, newLine)
	}

	area := mapOutVents(lines)
	elapsed := time.Since(start)
	utils.PostSolutionWithTime(5, 1, elapsed, area.Overlap())

	area = mapDiagonals(lines, area)
	elapsed = time.Since(start)

	utils.PostSolutionWithTime(5, 2, elapsed, area.Overlap())
}

func mapDiagonals(diagonals []Line, area hydrothermalVents) hydrothermalVents {
	for _, diagonal := range diagonals {
		if !diagonal.diagonal {
			continue
		}

		maxX, maxY := area.maxX, area.maxY
		xDirection := -1
		yDirection := -1
		if diagonal.StartX < diagonal.EndX {
			xDirection = 1
		}
		if diagonal.StartY < diagonal.EndY {
			yDirection = 1
		}
		currentX, currentY := diagonal.StartX, diagonal.StartY
		endX := diagonal.EndX
		// Because of 1 off errors (and I want to acutal check the last point aswell)
		endX += xDirection
		for currentX != endX {
			if maxX < currentX {
				fmt.Println("Increasing maxX")
				maxX = currentX
			}
			if maxY < currentY {
				fmt.Println("Increasing maxY")
				maxY = currentY
			}
			if _, ok := area.VisitedLocation[currentX]; !ok {
				area.VisitedLocation[currentX] = make(map[int]int)
			}

			area.VisitedLocation[currentX][currentY] = area.VisitedLocation[currentX][currentY] + 1
			currentX += xDirection
			currentY += yDirection
		}
	}

	return area
}

func mapOutVents(lines []Line) (area hydrothermalVents) {
	area.VisitedLocation = make(map[int]map[int]int)
	maxY, maxX := 0, 0
	for _, line := range lines {
		if line.diagonal {
			continue
		}
		for x := line.StartX; x <= line.EndX; x++ {
			if x > maxX {
				maxX = x
			}
			if _, ok := area.VisitedLocation[x]; !ok {
				area.VisitedLocation[x] = make(map[int]int)
			}

			for y := line.StartY; y <= line.EndY; y++ {
				if y > maxY {
					maxY = y
				}
				area.VisitedLocation[x][y] = area.VisitedLocation[x][y] + 1
			}
		}
	}

	area.maxX = maxX
	area.maxY = maxY
	return area
}

func parseLine(val string) Line {
	line := Line{}

	points := strings.Split(val, " -> ")
	for pointKey, point := range points {
		splittedPoint := strings.Split(point, ",")
		for key, coordinate := range splittedPoint {
			cor, err := strconv.Atoi(coordinate)
			if err != nil {
				panic(err)
			}
			if key == 0 {
				if pointKey == 0 {
					line.StartX = cor
				} else {
					line.EndX = cor
				}
			} else {
				if pointKey == 0 {
					line.StartY = cor
				} else {
					line.EndY = cor
				}
			}
		}
	}

	startX, startY, endX, endY := line.StartX, line.StartY, line.EndX, line.EndY
	if line.EndY != line.StartY && line.EndX != line.StartX {
		line.diagonal = true
	}
	if !line.diagonal {
		if startX > endX || startY > endY {
			line.StartX = endX
			line.EndX = startX
			line.StartY = endY
			line.EndY = startY
		}
	}

	return line
}
