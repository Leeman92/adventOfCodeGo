package day3

import (
	"fmt"
	"strings"
)

const tree = "#"

// Solve runs the puzzle program
func Solve(lines []string) {
	partOne(lines)
	partTwo(lines)
}

func partOne(lines []string) {
	treeCounter := 0
	for lineNumber, line := range lines {
		if !isTree(lineNumber, line, 3, 1) {
			continue
		}
		treeCounter++
	}

	fmt.Printf("You hit %d trees in part 1\n", treeCounter)
}

func partTwo(lines []string) {
	treeCounter := 0
	slopes := [5][2]int{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	for _, slope := range slopes {
		slopeCounter := 0
		for i := 0; i < len(lines); i += slope[1] {
			if !isTree(i, lines[i], slope[0], slope[1]) {
				continue
			}
			slopeCounter++
		}
		if treeCounter == 0 {
			treeCounter = slopeCounter
		} else {
			treeCounter = treeCounter * slopeCounter
		}
	}

	fmt.Printf("Answer for part 2 is: %d\n", treeCounter)
}

func isTree(lineNumber int, line string, right int, skip int) bool {
	lineSplice := strings.Split(line, "")
	index := 0

	if lineNumber > 0 {
		index = ((lineNumber / skip) * right) % len(lineSplice)
	}

	return checkValue(lineSplice[index])
}

func checkValue(value string) bool {
	return value == tree
}
