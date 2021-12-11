package utils

import (
	"fmt"
	"math/big"
	"regexp"
	"strconv"
	"strings"
	"time"
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

// ConvertLinesToIntSlice takes the line from the input and parses its values to an integer slice.
func ConvertLinesToIntSlice(lines []string) map[int]map[int]int {
	result := make(map[int]map[int]int)

	for outerKey, line := range lines {
		splitLine := strings.Split(line, "")
		for innerKey, val := range splitLine {
			_, ok := result[outerKey]
			if !ok {
				result[outerKey] = make(map[int]int)
			}
			numb, err := strconv.Atoi(val)
			if err != nil {
				panic(err)
			}
			result[outerKey][innerKey] = numb
		}
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

// UniqueIntSlice takes in a onedimensional slice and removes all duplicate entries
func UniqueIntSlice(intSlice []int) []int {
	keys := make(map[int]bool)
	list := []int{}
	for _, entry := range intSlice {
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
func PostSolutionWithTime(day int, part int, duration time.Duration, answer ...interface{}) {
	fmt.Printf("========== DAY %d -- PART %d ==========\n", day, part)
	fmt.Printf("          The answer is: %v\n", answer)
	fmt.Printf("          It took: %dns\n", duration)
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

func Remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

func BinarySliceToInt(binaryRepresantation []string) int {
	i, err := strconv.ParseInt(strings.Join(binaryRepresantation, ""), 2, 64)
	if err != nil {
		panic(err)
	}

	return int(i)
}

func GetRelevantBit(byteCollection []string, position int, least bool) string {
	byteCount := len(byteCollection)
	counter := 0

	for _, val := range byteCollection {
		byteSlice := strings.Split(val, "")
		bit, _ := strconv.Atoi(byteSlice[position])
		counter += bit
	}

	if !least {
		if byteCount%2 == 0 && counter == byteCount/2 {
			if counter >= (byteCount / 2) {
				return "1"
			} else {
				return "0"
			}
		} else {
			if counter > (byteCount / 2) {
				return "1"
			} else {
				return "0"
			}
		}
	} else {
		if byteCount%2 == 0 && counter == byteCount/2 {
			if counter >= (byteCount / 2) {
				return "0"
			} else {
				return "1"
			}
		} else {
			if counter > (byteCount / 2) {
				return "0"
			} else {
				return "1"
			}
		}

	}
}

func FilterSlice(haystack []string, needle string, position int) (cleanedHaystack []string) {
	for _, val := range haystack {
		valSlice := strings.Split(val, "")
		if valSlice[position] != needle {
			continue
		}
		cleanedHaystack = append(cleanedHaystack, val)
	}

	return cleanedHaystack
}

func ParseInputToIntSlice(input []string) (output []int) {
	for _, val := range input {
		valSlice := strings.Split(val, ",")
		for _, str := range valSlice {
			num, err := strconv.Atoi(str)
			if err != nil {
				panic(err)
			}
			output = append(output, num)
		}
	}

	return output
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
