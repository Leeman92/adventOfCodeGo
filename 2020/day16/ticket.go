package day16

import "github.com/l33m4n123/adventOfCodeGo/2020/utils"

type Ticket struct {
	values         []int
	possibleValues map[string][]int
	uniqueValues   map[string][]int
}

func (t *Ticket) checkRules(rules Rules) {
	if t.possibleValues == nil {
		t.possibleValues = make(map[string][]int)
	}
	for _, val := range t.values {
		for ruleName, ruleValues := range rules.rules {
			for index, value := range ruleValues {
				if value == val {
					t.possibleValues[ruleName] = append(t.possibleValues[ruleName], index)
				}
			}
		}
	}
}

func (t *Ticket) addUniqueValues(name string, newValues []int) {
	tmp := t.uniqueValues[name]
	for _, val := range newValues {
		tmp = append(tmp, val)
	}

	tmp = utils.UniqueIntSlice(tmp)
	t.uniqueValues[name] = tmp
}
