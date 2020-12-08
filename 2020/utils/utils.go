package utils

import (
	"fmt"
	"strconv"
)

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