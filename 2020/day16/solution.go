package day16

import (
	"fmt"
	"strings"

	"github.com/l33m4n123/adventOfCodeGo/2020/utils"
)

func Solve(input []string) {
	rules := Rules{}
	rules = parseRules(input, rules)

	solution := partOne(rules, input)

	utils.PostSolution(16, 1, solution)

	solution = partTwo(rules, input)

	utils.PostSolution(16, 1, solution)
}

func parseRules(input []string, rules Rules) Rules {
	for _, line := range input {
		rules.addRule(line)
	}
	return rules
}

func partOne(rules Rules, input []string) (result int) {
	skipping := true
	for _, line := range input {
		if skipping {
			if line == "nearby tickets:" {
				skipping = false
			}
			continue
		}

		result += rules.isInputValid(line)
	}

	return result
}

func partTwo(rules Rules, input []string) (result int) {
	skipping := true
	validTickets := []Ticket{}
	rulesVal := map[string][]int{}

	for _, line := range input {
		if skipping {
			if line == "nearby tickets:" {
				skipping = false
			}
			continue
		}

		if rules.isInputValid(line) > 0 {
			continue
		}

		splittedVal := strings.Split(line, ",")
		intVals := utils.ConvertLinesToInt(splittedVal)
		ticket := Ticket{}
		ticket.values = intVals
		ticket.checkRules(rules)

		validTickets = append(validTickets, ticket)
	}
	done := false
	counter := 0
	for !done {
		done = true
		counter++
		for idx, ticket := range validTickets {
			valFound := make(map[int]int)
			for _, val := range ticket.possibleValues {
				for _, v := range val {
					valFound[v]++
				}
			}

			for name, val := range ticket.possibleValues {
				for _, v := range val {
					if valFound[v] == 1 {
						if ticket.uniqueValues == nil {
							ticket.uniqueValues = make(map[string][]int)
						}
						tmp := ticket.uniqueValues[name]
						tmp = append(tmp, v)
						ticket.addUniqueValues(name, tmp)
					}
				}
			}

			validTickets[idx] = ticket
			for _, val := range ticket.uniqueValues {
				if len(val) > 1 {
					done = false
				}
			}
		}

		if counter == 100 {
			done = true
		}
	}

	for _, ticket := range validTickets {
		for name, val := range ticket.uniqueValues {
			tmp := rulesVal[name]
			for _, v := range val {
				tmp = append(tmp, v)
			}
			tmp = utils.UniqueIntSlice(tmp)
			rulesVal[name] = tmp
		}
	}

	valRemove := []int{}
	for i := 0; i < 500; i++ {
		foundVal := map[int]int{}
		for _, valueToRemove := range valRemove {
			for index, val := range rulesVal {
				for idx, v := range val {
					if v == valueToRemove {
						rulesVal[index] = utils.Remove(rulesVal[index], idx)
					}
				}
			}
		}
		for _, val := range rulesVal {
			for _, v := range val {
				foundVal[v]++
			}
		}

		newRules := map[string][]int{}
		for name, val := range rulesVal {
			for _, v := range val {
				if foundVal[v] == 1 {
					tmpVal := newRules[name]
					tmpVal = append(tmpVal, v)
					tmpVal = utils.UniqueIntSlice(tmpVal)
					newRules[name] = tmpVal
				}
			}
		}

		for name, val := range newRules {
			if len(val) == 1 {
				valRemove = append(valRemove, val[0])
			}
			rulesVal[name] = val
		}
	}

	for name, val := range rulesVal {
		fmt.Printf("%v -- %v\n", name, val)
	}
	return result
}
