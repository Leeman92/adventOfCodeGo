package utils

import "strconv"

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
