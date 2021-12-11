package day11

import (
	"fmt"
	"strconv"
	"strings"
)

type Octopus struct {
	EnergyLevel int
	Flashed     bool
}

func Solve(lines []string) {
	playingField := parseLines(lines)
	fmt.Println(playingField)
}

func parseLines(lines []string) (playingField map[int]map[int]Octopus) {
	for row, line := range lines {
		_, ok := playingField[row]
		if !ok {
			playingField[row] = make(map[int]Octopus)
		}
		splitLine := strings.Split(line, "")
		for column, entry := range splitLine {
			numb, err := strconv.Atoi(entry)
			if err != nil {
				panic(err)
			}
			playingField[row][column] = Octopus{EnergyLevel: numb, Flashed: false}
		}
	}

	return playingField
}
