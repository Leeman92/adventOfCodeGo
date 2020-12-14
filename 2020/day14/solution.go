package day14

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/l33m4n123/adventOfCodeGo/2020/utils"
)

func Solve(input []string) {
	program := setupProgramms(input)

	totalAnswer := runAllProgramsPartOne(program)
	utils.PostSolution(14, 1, totalAnswer)

	totalAnswerPartTwo := runAllProgramsPartTwo(program)
	utils.PostSolution(14, 2, totalAnswerPartTwo)
}

func runAllProgramsPartOne(programs []Program) (result int) {
	cache := make(map[int]int)
	for i := 0; i < len(programs); i++ {
		program := programs[i]
		cache = program.Run(cache)
	}

	for _, val := range cache {
		result += val
	}

	return result
}

func runAllProgramsPartTwo(programs []Program) (result int) {
	cache := make(map[int]int)
	for i := 0; i < len(programs); i++ {
		program := programs[i]
		cache = program.RunPart2(cache)
	}

	for _, val := range cache {
		result += val
	}

	return result
}

func setupProgramms(input []string) []Program {
	programs := []Program{}
	var program Program

	maskPattern := regexp.MustCompile("mask = (?P<mask>.*)")
	operatorPattern := regexp.MustCompile("mem\\[(?P<memoryLocation>\\d*)\\] = (?P<value>\\d*)")

	for index, line := range input {
		maskMatches := utils.FindStringSubmatchWithNamedMatches(maskPattern, line)
		if len(maskMatches) > 0 {
			if index > 0 {
				programs = append(programs, program)
			}
			program = Program{}
			program.operators = []map[string]int{}

			mask := maskMatches["mask"]
			andMask := strings.ReplaceAll(mask, "X", "1")
			orMask := strings.NewReplacer("X", "0").Replace(mask)
			noChange := strings.NewReplacer("X", "0", "1", "0", "0", "1").Replace(mask)
			floatMask := getFloatMasks(mask)

			andMaskInt64, andParseError := strconv.ParseInt(andMask, 2, 64)
			if andParseError != nil {
				panic(andParseError)
			}

			orMaskInt64, orParseError := strconv.ParseInt(orMask, 2, 64)
			if orParseError != nil {
				panic(orParseError)
			}

			noChangeInt64, noChangeParseError := strconv.ParseInt(noChange, 2, 64)
			if noChangeParseError != nil {
				panic(noChangeParseError)
			}

			program.andMask = int(andMaskInt64)
			program.orMask = int(orMaskInt64)
			program.floatMask = floatMask
			program.noChangeMask = int(noChangeInt64)
			continue
		}

		operatorMatches := utils.FindStringSubmatchWithNamedMatches(operatorPattern, line)

		if len(operatorMatches) <= 0 {
			continue
		}

		memLocation, memError := strconv.ParseInt(operatorMatches["memoryLocation"], 10, 64)
		if memError != nil {
			panic(memError)
		}

		value, valError := strconv.ParseInt(operatorMatches["value"], 10, 64)
		if valError != nil {
			panic(valError)
		}
		tmp := make(map[string]int)
		tmp["memLocation"] = int(memLocation)
		tmp["value"] = int(value)
		program.operators = append(program.operators, tmp)
	}

	programs = append(programs, program)

	return programs
}

// 1X0X => [0000, 0001, 0100, 0101] => [0, 1, 4, 5]
func getFloatMasks(s string) []int {
	a := []int{}
	for idx, char := range s {
		if char == 'X' {
			a = append(a, 1<<(len(s)-idx-1))
		}
	}

	b := []int{0}
	for _, i := range a {
		for _, j := range b {
			b = append(b, i|j)
		}
	}

	return b
}
