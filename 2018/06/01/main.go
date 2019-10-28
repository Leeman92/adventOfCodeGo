package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
)

// PointsInSpaceMap - keeps track of all the points in the world
var PointsInSpaceMap map[int]map[int]PointsInSpace

// FilledPointsInSpace - keeps track of the relevant Points
var FilledPointsInSpace []PointsInSpace

var controlledArea map[PointsInSpace]int

// Coordinates - represantation of coordinates in WorldSpace
type Coordinates struct {
	x, y int
}

// PointsInSpace - struct for easier working
type PointsInSpace struct {
	coordinates Coordinates
	display     rune
	state       bool
}

func main() {
	fmt.Println("=========== START ==========")
	PointsInSpaceMap = make(map[int]map[int]PointsInSpace)
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(b), "\n")

	maxX, maxY := calculateWorldHeight(lines)

	fmt.Printf("World is %d, %d big\n", maxX, maxY)
	setupWorldMap(maxX, maxY)
	fillPointsInSpace(lines)
	fillMap(maxX, maxY)
	getBiggestArea()

	if maxX <= 50 && maxY <= 12 {
		drawMap(maxX, maxY)
	}
}

func calculateWorldHeight(lines []string) (int, int) {
	maxX := 0
	maxY := 0
	for _, pointsInSpace := range lines {
		currX, currY := getCoordinatesFromLine(pointsInSpace)
		if maxX < currX {
			maxX = currX
		}
		if maxY < currY {
			maxY = currY
		}
	}

	// Increase Mapsize for drawing and stuff
	maxX++
	maxY++

	return maxX, maxY
}

func setupWorldMap(maxX int, maxY int) {
	for x := 0; x <= maxX; x++ {
		PointsInSpaceMap[x] = make(map[int]PointsInSpace)
		for y := 0; y <= maxY; y++ {
			PointsInSpaceMap[x][y] = PointsInSpace{Coordinates{x, y}, '.', false}
		}
	}
}

func fillPointsInSpace(lines []string) {
	for pos, pointsInSpace := range lines {
		currX, currY := getCoordinatesFromLine(pointsInSpace)
		currentPoint := PointsInSpace{Coordinates{currX, currY}, toChar(pos + 1), true}
		PointsInSpaceMap[currX][currY] = currentPoint
		FilledPointsInSpace = append(FilledPointsInSpace, currentPoint)
	}
}

func fillMap(maxX int, maxY int) {
	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			shortestDistance := math.MaxInt64
			amountOfPointsAtShortestDistance := 0
			var pointOfShortestDistance PointsInSpace
			for _, pointInSpace := range FilledPointsInSpace {
				if PointsInSpaceMap[x][y].state {
					continue
				}
				distance := getEuclideanDistance(pointInSpace.coordinates, Coordinates{x, y})
				if shortestDistance > distance {
					shortestDistance = distance
					amountOfPointsAtShortestDistance = 1
					pointOfShortestDistance = pointInSpace
				} else if shortestDistance == distance {
					amountOfPointsAtShortestDistance++
				}
			}

			if shortestDistance > 0 && amountOfPointsAtShortestDistance == 1 {
				displayString := string(pointOfShortestDistance.display)
				displayString = strings.ToLower(displayString)
				displayRune := []rune(displayString)

				PointsInSpaceMap[x][y] = PointsInSpace{Coordinates{x, y}, rune(displayRune[0]), false}
				if _, ok := controlledArea[pointOfShortestDistance]; !ok {
					if controlledArea == nil {
						controlledArea = make(map[PointsInSpace]int)
					}
					controlledArea[pointOfShortestDistance] = 1
				}

				if controlledArea[pointOfShortestDistance] < math.MaxInt64 {
					controlledArea[pointOfShortestDistance]++
				}

				if atBorder(Coordinates{x, y}, maxX, maxY) {
					controlledArea[pointOfShortestDistance] = math.MaxInt64
				}
			}
		}
	}
}

func getBiggestArea() {
	biggestNonFiniteArea := 0
	for _, areaSize := range controlledArea {
		if areaSize > biggestNonFiniteArea && areaSize < math.MaxInt64 {
			biggestNonFiniteArea = areaSize
		}
	}

	fmt.Printf("Biggest nonfinite Area encloses %d tiles\n", biggestNonFiniteArea)
}

func drawMap(maxX int, maxY int) {
	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			characterToPrint := PointsInSpaceMap[x][y].display
			fmt.Printf("%v ", string(characterToPrint))
		}
		fmt.Printf("\n")
	}
}

// Helper funcs
func getCoordinatesFromLine(line string) (int, int) {
	coordinates := strings.Split(line, ", ")
	currX, _ := strconv.Atoi(strings.TrimSpace(coordinates[0]))
	currY, _ := strconv.Atoi(strings.TrimSpace(coordinates[1]))

	return currX, currY
}

// Abs returns the absolute value of x.
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func toChar(i int) rune {
	return rune('A' - 1 + i)
}

func toCharStr(i int) string {
	return string('A' - 1 + i)
}

func getEuclideanDistance(firstCoordinate Coordinates, secondCoordinate Coordinates) int {
	totalDistance := 0
	totalDistance += Abs(firstCoordinate.x - secondCoordinate.x)
	totalDistance += Abs(firstCoordinate.y - secondCoordinate.y)
	//fmt.Printf("Distance between %v and %v is %d\n", firstCoordinate, secondCoordinate, totalDistance)
	return totalDistance
}

func atBorder(pointInSpace Coordinates, maxX int, maxY int) bool {
	currX := pointInSpace.x
	currY := pointInSpace.y

	returnValue := false
	// Top Border
	if currY == 0 {
		returnValue = true
	}

	// Right Border
	if currX == maxX {
		returnValue = true
	}

	// Top Border
	if currY == maxY {
		returnValue = true
	}

	// Right Border
	if currX == 0 {
		returnValue = true
	}

	//fmt.Printf("%d, %d -- %v\n", currX, currY, returnValue)

	return returnValue
}
