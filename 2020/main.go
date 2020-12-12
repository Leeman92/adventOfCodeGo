package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/l33m4n123/adventOfCodeGo/2020/day1"
	"github.com/l33m4n123/adventOfCodeGo/2020/day10"
	"github.com/l33m4n123/adventOfCodeGo/2020/day11"
	"github.com/l33m4n123/adventOfCodeGo/2020/day12"
	"github.com/l33m4n123/adventOfCodeGo/2020/day2"
	"github.com/l33m4n123/adventOfCodeGo/2020/day3"
	"github.com/l33m4n123/adventOfCodeGo/2020/day4"
	"github.com/l33m4n123/adventOfCodeGo/2020/day5"
	"github.com/l33m4n123/adventOfCodeGo/2020/day6"
	"github.com/l33m4n123/adventOfCodeGo/2020/day7"
	"github.com/l33m4n123/adventOfCodeGo/2020/day8"
	"github.com/l33m4n123/adventOfCodeGo/2020/day9"
)

const realInputFileName = "input.txt"
const testInputFileName = "input.test.txt"

func main() {
	status := make(chan bool)

	day := flag.Int("day", 1, "Tells the program what day to run")
	test := flag.Bool("test", false, "Running the tests only?")
	flag.Parse()
	fmt.Printf("Start the program for day %d. Running tests: %v\n", *day, *test)

	fileName := fmt.Sprintf("day%d/%s", *day, realInputFileName)
	if *test {
		fileName = fmt.Sprintf("day%d/%s", *day, testInputFileName)
	}

	lines, err := readLines(fileName)
	if err != nil {
		panic(err)
	}

	go runSolution(*day, lines, status, *test)
	<-status
}

func runSolution(day int, lines []string, status chan bool, test bool) {
	fmt.Println("Running solution")
	switch day {
	case 1:
		day1.Solve(lines)
	case 2:
		day2.Solve(lines)
	case 3:
		day3.Solve(lines)
	case 4:
		day4.Solve(lines)
	case 5:
		day5.Solve(lines)
	case 6:
		day6.Solve(lines)
	case 7:
		day7.Solve(lines)
	case 8:
		day8.Solve(lines)
	case 9:
		day9.Solve(lines, test)
	case 10:
		day10.Solve(lines)
	case 11:
		day11.Solve(lines)
	case 12:
		day12.Solve(lines)
	default:
		panic(fmt.Sprintf("No solution for day %d implemented yet!\n", day))
	}
	status <- true
}

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
