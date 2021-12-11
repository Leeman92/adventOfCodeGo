package day9

import (
	"fmt"
	"github.com/l33m4n123/adventOfCodeGo/2021/utils"
	"sort"
)

func Solve(lines []string) {
	coordinateSystem := utils.ConvertLinesToIntSlice(lines)
	maxY := len(coordinateSystem)
	maxX := len(coordinateSystem[0])

	riskLevel := 0
	lowPoints := make(map[int]map[int]int)
	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			value := coordinateSystem[y][x]
			neighbours := getNeighbourValue(y, x, coordinateSystem)
			lowest := true
			for _, neighbour := range neighbours {
				if neighbour <= value {
					lowest = false
				}
			}

			if lowest {
				_, ok := lowPoints[y]
				if !ok {
					lowPoints[y] = make(map[int]int)
				}
				lowPoints[y][x] = value
				riskLevel += 1 + value
			}
		}
	}

	utils.PostSolution(9, 1, riskLevel)

	var visitedLocations []string
	var basins []int
	//
	//sort.Ints(basins)
	//
	//fmt.Println(basins, basins[len(basins)-1]*basins[len(basins)-2]*basins[len(basins)-3])

	for y, row := range lowPoints {
		for x, _ := range row {
			res, _ := getBasinCount(y, x, coordinateSystem, visitedLocations)
			basins = append(basins, res)
		}
	}
	sort.Ints(basins)

	utils.PostSolution(9, 2, basins[len(basins)-1]*basins[len(basins)-2]*basins[len(basins)-3])

}

func getBasinCount(y int, x int, system map[int]map[int]int, locations []string) (int, []string) {
	basinCount := 0
	if alreadyVisited(y, x, locations) || system[y][x] == 9 {
		return basinCount, locations
	}

	locations = append(locations, fmt.Sprintf("%d|%d", y, x))
	basinCount++

	// Only non diagonal neighbours
	x1 := x + 1
	x2 := x - 1
	y1 := y + 1
	y2 := y - 1

	_, ok1 := system[y1][x]
	_, ok2 := system[y2][x]
	_, ok3 := system[y][x1]
	_, ok4 := system[y][x2]

	count := 0
	if ok1 {
		count, locations = getBasinCount(y1, x, system, locations)
		basinCount += count
	}
	if ok2 {
		count, locations = getBasinCount(y2, x, system, locations)
		basinCount += count
	}
	if ok3 {
		count, locations = getBasinCount(y, x1, system, locations)
		basinCount += count
	}
	if ok4 {
		count, locations = getBasinCount(y, x2, system, locations)
		basinCount += count
	}

	return basinCount, locations
}

func alreadyVisited(y int, x int, locations []string) bool {
	for _, location := range locations {
		if location == fmt.Sprintf("%d|%d", y, x) {
			return true
		}
	}
	return false
}

func getNeighbourValue(y int, x int, coordinateSystem map[int]map[int]int) []int {
	var neighbours []int
	// Only non diagonal neighbours
	x1 := x + 1
	x2 := x - 1
	y1 := y + 1
	y2 := y - 1

	value1, ok1 := coordinateSystem[y1][x]
	value2, ok2 := coordinateSystem[y2][x]
	value3, ok3 := coordinateSystem[y][x1]
	value4, ok4 := coordinateSystem[y][x2]

	if ok1 {
		neighbours = append(neighbours, value1)
	}
	if ok2 {
		neighbours = append(neighbours, value2)
	}
	if ok3 {
		neighbours = append(neighbours, value3)
	}
	if ok4 {
		neighbours = append(neighbours, value4)
	}

	return neighbours
}
