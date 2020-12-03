package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	status := make(chan bool)
	lines, err := readLines("input.txt")

	if err != nil {
		panic(err)
	}

	go partOne(status, lines)
	<-status
	go partTwo(status, lines)
	<-status
}

func partOne(status chan bool, lines []string) {
	treeCounter := 0
	for lineNumber, line := range lines {
		if !isTree(lineNumber, line, 3, 1) {
			continue
		}
		treeCounter++
	}

	fmt.Printf("You hit %d trees in part 1\n", treeCounter)
	status <- true
}

func partTwo(status chan bool, lines []string) {
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
	status <- true
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
	return value == "#"
}

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineStr := scanner.Text()
		lines = append(lines, lineStr)
	}

	return lines, scanner.Err()
}
