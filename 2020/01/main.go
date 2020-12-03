package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	status := make(chan bool)
	go partOne(status)
	<-status
	go partTwo(status)
	<-status
}

func partOne(status chan bool) {
	lines, err := readLines("input.txt")

	if err != nil {
		panic(err)
	}

	for outerKey, x := range lines {
		for innerKey, y := range lines {
			if innerKey == outerKey || innerKey <= outerKey {
				continue
			}

			if x+y == 2020 {
				fmt.Printf("%d + %d = 2020\n", x, y)
				fmt.Printf("Your answer for Part 1 is %d\n", x*y)
				status <- true
				return
			}
		}
	}

	status <- true
}

func partTwo(status chan bool) {
	lines, err := readLines("input.txt")

	if err != nil {
		panic(err)
	}

	for xKey, x := range lines {
		for yKey, y := range lines {
			for zKey, z := range lines {
				if xKey == yKey ||
					xKey == zKey ||
					yKey == zKey ||
					yKey < xKey ||
					zKey < yKey {
					continue
				}

				if x+y+z == 2020 {
					fmt.Printf("%d + %d + %d = 2020\n", x, y, z)
					fmt.Printf("Your answer for Part 2 is %d\n", x*y*z)
					status <- true
					return
				}
			}
		}
	}

	status <- true
}

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineStr := scanner.Text()
		lineInt, castErr := strconv.Atoi(lineStr)
		if castErr != nil {
			return nil, err
		}
		lines = append(lines, lineInt)
	}

	return lines, scanner.Err()
}
