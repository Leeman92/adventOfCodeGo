package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"time"
)

func main() {
	now := time.Now()
	globalStart := now.UnixNano()
	fmt.Println("=========== START ==========")
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(b), "\n")
	shortestLength := 0
	shortestPolymer := ""
	removedUnit := ""
	for _, polymer := range lines {
		shortestLength = len(polymer)
		allUnitTypes := getPolymerUnits(polymer)
		for _, unitType := range allUnitTypes {
			leftOverPolymer := removeOccurenceFromPolymer(unitType, polymer)
			leftOverPolymer = reactPolymer(leftOverPolymer)
			if shortestLength > len(leftOverPolymer) {
				shortestLength = len(leftOverPolymer)
				shortestPolymer = leftOverPolymer
				removedUnit = unitType
			}
		}
	}

	printSolution(shortestPolymer, removedUnit)
	now = time.Now()
	globalEnd := now.UnixNano()
	fmt.Printf("Took %v total nano seconds", globalEnd-globalStart)
}

func removeOccurenceFromPolymer(unitType string, polymer string) string {
	polymer = strings.ReplaceAll(polymer, strings.ToLower(unitType), "")
	polymer = strings.ReplaceAll(polymer, strings.ToUpper(unitType), "")
	return polymer
}

func getPolymerUnits(polymer string) []string {
	polymerSlice := strings.Split(polymer, "")
	var unitsInPolymer []string

	for _, currentUnit := range polymerSlice {
		if contains(unitsInPolymer, strings.ToLower(currentUnit)) {
			continue
		}

		unitsInPolymer = append(unitsInPolymer, strings.ToLower(currentUnit))
	}

	return unitsInPolymer
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func printSolution(polymer string, unitType string) {
	polymer = strings.TrimSpace(polymer)
	fmt.Printf("========= SOLUTION =========\nThe leftover polymer is %d units long after removing %s/%s units.\n============================\n", len(polymer), strings.ToLower(unitType), strings.ToUpper(unitType))
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
				toBeReplaced := strings.ToLower(lastUnit) + strings.ToUpper(lastUnit)
				toBeReplacedInversed := strings.ToUpper(lastUnit) + strings.ToLower(lastUnit)
				polymer = strings.ReplaceAll(polymer, toBeReplaced, "")
				polymer = strings.ReplaceAll(polymer, toBeReplacedInversed, "")
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
