package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	fmt.Println("=========== START ==========")
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(b), "\n")
	for _, polymer := range lines {
		leftOverPolymer := reactPolymer(polymer)
		printSolution(leftOverPolymer)
	}
}

func printSolution(polymer string) {
	polymer = strings.TrimSpace(polymer)
	fmt.Printf("========= SOLUTION =========\nThe leftover polymer is %d units long.\n============================", len(polymer))
}

func reactPolymer(polymer string) string {
	var lastUnit string
	var reacted bool
	var polymerSlice []string
	polymerSlice = strings.Split(polymer, "")

	for pos, currentUnit := range polymerSlice {
		if pos == 0 {
			lastUnit = currentUnit
			continue
		}
		if strings.ToLower(lastUnit) == strings.ToLower(currentUnit) {
			if lastUnit != currentUnit {
				reacted = true
				toBeReplaced := lastUnit + currentUnit
				polymer = strings.ReplaceAll(polymer, toBeReplaced, "")
				break
			}
		}

		lastUnit = currentUnit
	}

	if !reacted {
		return polymer
	}

	polymer = reactPolymer(polymer)
	return polymer
}
