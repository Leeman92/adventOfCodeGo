package day14

import (
	"fmt"
	"regexp"

	"github.com/l33m4n123/adventOfCodeGo/2020/utils"
)

type Program struct {
	andMask   uint64
	orMask    uint64
	operators map[int]int
}

func Solve(input []string) {
	program := setupProgramms(input)

	fmt.Println(program[0].andMask)
}

func setupProgramms(input []string) []*Program {
	programs := []*Program{}
	program := Program{}

	maskPattern := regexp.MustCompile("mask = (?P<mask>.*)")
	operatorPattern := regexp.MustCompile("mem\\[(?P<memoryLocation>\\d*)\\] = (?P<value>\\d*)")

	for index, line := range input {
		matchesMask := utils.FindStringSubmatchWithNamedMatches(maskPattern, line)
		if len(matchesMask) > 0 {
			if index > 0 {
				programs = append(programs, &program)
				program = Program{}
			}

			fmt.Println(matchesMask)

			program.andMask = 5
			program.orMask = 100
			continue
		}

		operatorMatches := utils.FindStringSubmatchWithNamedMatches(operatorPattern, line)

		if len(operatorMatches) <= 0 {
			continue
		}

		fmt.Println(operatorMatches)
	}

	programs = append(programs, &program)

	return programs
}
