package day7

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/l33m4n123/adventOfCodeGo/utils"
)

type bag struct {
	color       string
	canHoldBags map[string]int
	sumOfBags   int
}

func (b *bag) addBags(sum string) {
	i, _ := strconv.Atoi(sum)
	b.sumOfBags += i
}

func (b *bag) addHoldedBags(name string, amount int) {
	if b.canHoldBags == nil {
		b.canHoldBags = make(map[string]int)
	}
	b.canHoldBags[name] = amount
}

// Solve runs the puzzle
func Solve(input []string) {
	bags := parseBags(input)
	partOne(bags)

	partTwo(bags)
}

func partOne(bags []bag) {
	solution := findBagsThatHold("shiny gold", bags)

	utils.PostSolution(7, 1, len(solution))
}

func partTwo(bags []bag) {
	allBags := bagsHoldBy("shiny gold", bags)
	startBag := allBags["shiny gold"]
	solution := recursiveCounting("shiny gold", startBag, allBags)
	// For now I calc it myself D:
	utils.PostSolution(7, 2, solution)
}

func recursiveCounting(bagName string, bag map[string]int, allBags map[string]map[string]int) int {
	sol := 1

	for name, value := range bag {
		if name == "" {
			return 1
		}
		res := recursiveCounting(name, allBags[name], allBags)
		sol += value * res
	}

	if bagName == "shiny gold" {
		sol--
	}

	return sol
}

func bagsHoldBy(name string, bags []bag) map[string]map[string]int {
	totalBags := make(map[string]map[string]int)

	for _, bag := range bags {
		if bag.color != name {
			continue
		}

		totalBags[name] = bag.canHoldBags
		for color := range bag.canHoldBags {
			if color == "" {
				continue
			}

			result := bagsHoldBy(color, bags)
			for key, res := range result {
				totalBags[key] = res
			}
		}
	}

	return totalBags
}

func findBagsThatHold(name string, bags []bag) []string {
	possibleColors := []string{}
	addedNewStuff := false
	for _, bag := range bags {
		for color := range bag.canHoldBags {
			if color != name {
				continue
			}
			possibleColors = append(possibleColors, bag.color)
			addedNewStuff = true
		}
	}

	if addedNewStuff {
		for _, color := range possibleColors {
			result := findBagsThatHold(color, bags)
			for _, newColor := range result {
				possibleColors = append(possibleColors, newColor)
			}
		}
	}

	possibleColors = utils.UniqueStringSlice(possibleColors)

	return possibleColors
}

func parseBags(input []string) []bag {
	bags := []bag{}

	mainPattern := regexp.MustCompile(`(?P<selfColor>.*?) bags{0,1} contain ((?P<firstAmount>\d*) (?P<firstColor>.*?) bags{0,1})*`)
	secondaryPattern := regexp.MustCompile(`(?P<secondaryAmount>\d*?) (?P<secondaryColor>.*?) bags{0,1}`)
	for _, line := range input {
		b := bag{}
		splitInput := strings.Split(line, ",")
		for i, split := range splitInput {
			split = strings.TrimSpace(split)
			pattern := mainPattern
			if i == 0 {
				match := pattern.FindStringSubmatch(split)
				if len(match) == 0 {
					panic("No matches found!")
				}
				result := make(map[string]string)
				for i, name := range pattern.SubexpNames() {
					if i != 0 && name != "" {
						result[name] = match[i]
					}
				}
				b.color = result["selfColor"]
				res, convErr := strconv.Atoi(result["firstAmount"])
				if convErr != nil {
					res = 0
				}
				b.addHoldedBags(result["firstColor"], res)
				b.addBags(result["firstAmount"])
			} else {
				pattern = secondaryPattern
				match := pattern.FindStringSubmatch(split)
				result := make(map[string]string)
				for i, name := range pattern.SubexpNames() {
					if i != 0 && name != "" {
						result[name] = match[i]
					}
				}
				res, convErr := strconv.Atoi(result["secondaryAmount"])
				if convErr != nil {
					res = 0
				}
				b.addHoldedBags(result["secondaryColor"], res)
				b.addBags(result["secondaryAmount"])
			}
		}

		bags = append(bags, b)
	}

	return bags
}
