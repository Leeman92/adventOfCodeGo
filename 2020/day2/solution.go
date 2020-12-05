package day2

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// Solve runs the solution for the code
func Solve(lines []string) {
	partOne(lines)
	partTwo(lines)
}

func partOne(lines []string) {
	validCounter := 0
	for _, line := range lines {
		if !isValidOne(line) {
			continue
		}
		validCounter++
	}

	fmt.Printf("%d passwords for Part 1 were valid\n", validCounter)
}

func partTwo(lines []string) {
	validCounter := 0
	for _, line := range lines {
		if !isValidTwo(line) {
			continue
		}
		validCounter++
	}

	fmt.Printf("%d passwords for Part 2 were valid\n", validCounter)
}

func isValidOne(line string) bool {
	parts := strings.Split(line, ":")
	password := parts[1]
	rest := strings.Split(parts[0], " ")
	amount := strings.Split(rest[0], "-")
	min := amount[0]
	max := amount[1]
	pattern := rest[1]

	expression, err := regexp.Compile(fmt.Sprintf("^([^%s]*%s[^%s]*){%s,%s}$", pattern, pattern, pattern, min, max))
	if err != nil {
		panic(err)
	}

	return expression.Match([]byte(password))
}

func isValidTwo(line string) bool {
	parts := strings.Split(line, ":")
	password := strings.TrimSpace(parts[1])
	passwordSlice := strings.Split(password, "")
	rest := strings.Split(parts[0], " ")
	positions := strings.Split(rest[0], "-")
	pos1, _ := strconv.Atoi(positions[0])
	pos2, _ := strconv.Atoi(positions[1])
	pattern := rest[1]

	if passwordSlice[pos1-1] == pattern && passwordSlice[pos2-1] == pattern {
		return false
	}
	if passwordSlice[pos1-1] == pattern || passwordSlice[pos2-1] == pattern {
		return true
	}

	return false
}
