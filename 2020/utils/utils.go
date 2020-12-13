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

// FindStringSubmatchWithNamedMatches is a wrapper function to match a pattern in a string and return the named capture groups
func FindStringSubmatchWithNamedMatches(pattern *regexp.Regexp, input string) map[string]string {
	match := pattern.FindStringSubmatch(input)
	if len(match) == 0 {
		panic("No matches found!")
	}
	result := make(map[string]string)
	for i, name := range pattern.SubexpNames() {
		if i != 0 && name != "" {
			result[name] = match[i]
		}
	}

	return result
}

// GetGCD returns the greatest common divisor
func GetGCD(a, b int64) int64 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// ModInverse returns the inverse of a modulo m
func ModInverse(a, b int64) int64 {
	// fact if b is prime. then a*a * (b-1) == 1 for any a
	// => a*a*(b-2)*a == 1
	// so a*a*(b-2) is a modular inverse of a!
	return modPow(a, b-2, b)
}

func modPow(base, exponent, mod int64) int64 {
	if exponent == 0 {
		return 1
	} else if exponent%2 == 0 {
		return modPow((base*base)%mod, exponent/2, mod)
	} else {
		return (base * modPow(base, exponent-1, mod)) % mod
	}
}
