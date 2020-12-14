package utils

import (
	"fmt"
	"math/big"
	"regexp"
	"strconv"
)

var one = big.NewInt(1)

type Coordinates struct {
	X, Y int
}

// ConvertLinesToInt takes the line from the input and parses its values to integer.
func ConvertLinesToInt(lines []string) []int {
	result := []int{}

	for _, line := range lines {
		intLine, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		result = append(result, intLine)
	}

	return result
}

// UniqueStringSlice takes in a onedimensional slice and removes all duplicate entries
func UniqueStringSlice(stringSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range stringSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

// UniqueMultipleStringSlice takes in a 2dimensional slice and makes sure each internal slice is unique. Not caring if a slice is same as another one
func UniqueMultipleStringSlice(slice [][]string) [][]string {
	list := make([][]string, len(slice))
	for key, internalSlice := range slice {
		list[key] = UniqueStringSlice(internalSlice)
	}

	return list
}

// PostSolution is a wrapper to post the result of the
func PostSolution(day int, part int, answer ...interface{}) {
	fmt.Printf("========== DAY %d -- PART %d ==========\n", day, part)
	fmt.Printf("          The answer is: %v\n", answer)
	fmt.Printf("=====================================\n\n")
}

// PrintDebug prints a debug message. Used alot for me during testing x)
func PrintDebug(message string, values ...interface{}) {
	fmt.Printf("========== DEBUG ==========\n")
	fmt.Printf(message, values...)
	fmt.Println("")
	fmt.Printf("=====================================\n\n")
}

// FindStringSubmatchWithNamedMatches is a wrapper function to match a pattern in a string and return the named capture groups
func FindStringSubmatchWithNamedMatches(pattern *regexp.Regexp, input string) map[string]string {
	match := pattern.FindStringSubmatch(input)
	result := make(map[string]string)
	if len(match) == 0 {
		return result
	}

	for i, name := range pattern.SubexpNames() {
		if i != 0 && name != "" {
			result[name] = match[i]
		}
	}

	return result
}
