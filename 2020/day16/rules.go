package day16

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/l33m4n123/adventOfCodeGo/2020/utils"
)

// Rules is a struct that keeps track of the validValues
type Rules struct {
	rules map[string][]int
}

// SetRule allows me to set the specifics of the rule
func (r *Rules) addRule(line string) {
	rulePattern := regexp.MustCompile("(?P<ruleName>.*): (?P<LowerLimitOne>\\d{1,})-(?P<UpperLimitOne>\\d{1,}) or (?P<lowerLimitTwo>\\d{1,})-(?P<UpperLimitTwo>\\d{1,})")

	matches := utils.FindStringSubmatchWithNamedMatches(rulePattern, line)
	if len(matches) == 0 {
		return
	}

	if r.rules == nil {
		r.rules = make(map[string][]int)
	}

	validNumbers := []int{}

	lowerLimitOne, lowerLimitOneErr := strconv.Atoi(matches["LowerLimitOne"])
	if lowerLimitOneErr != nil {
		panic(lowerLimitOneErr)
	}

	upperLimitOne, upperLimitOneErr := strconv.Atoi(matches["UpperLimitOne"])
	if upperLimitOneErr != nil {
		panic(upperLimitOneErr)
	}

	lowerLimitTwo, lowerLimitTwoErr := strconv.Atoi(matches["lowerLimitTwo"])
	if lowerLimitTwoErr != nil {
		panic(lowerLimitTwoErr)
	}

	upperLimitTwo, upperLimitTwoErr := strconv.Atoi(matches["UpperLimitTwo"])
	if upperLimitTwoErr != nil {
		panic(upperLimitTwoErr)
	}

	for i := lowerLimitOne; i <= upperLimitOne; i++ {
		validNumbers = append(validNumbers, i)
	}

	for i := lowerLimitTwo; i <= upperLimitTwo; i++ {
		validNumbers = append(validNumbers, i)
	}

	r.rules[matches["ruleName"]] = validNumbers
}

func (r *Rules) isInputValid(input string) int {
	result := false
	inputSplit := strings.Split(input, ",")
	for _, split := range inputSplit {
		result = false
		splitVal, err := strconv.Atoi(split)
		if err != nil {
			return 0
		}
		for _, validNumbers := range r.rules {
			for _, number := range validNumbers {
				if splitVal == number {
					result = true
				}
			}
		}

		if !result {
			return splitVal
		}
	}
	return 0
}
