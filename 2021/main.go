package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/l33m4n123/adventOfCodeGo/2021/day1"
	"github.com/l33m4n123/adventOfCodeGo/2021/day2"
	"github.com/l33m4n123/adventOfCodeGo/2021/day3"
	"github.com/l33m4n123/adventOfCodeGo/2021/day4"
	"github.com/l33m4n123/adventOfCodeGo/2021/day5"
	"github.com/l33m4n123/adventOfCodeGo/2021/day6"
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
		break
	case 2:
		day2.Solve(lines)
		break
	case 3:
		day3.Solve(lines)
		break
	case 4:
		day4.Solve(lines)
		break
	case 5:
		day5.Solve(lines)
	case 6:
		day6.Solve(lines)
		break
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
