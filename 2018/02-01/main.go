package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// How many boxes fit rule 1, how many rule 2
	var rule1, rule2 int
	s := bufio.NewScanner(f)
	for s.Scan() {
		var boxID string
		_, err := fmt.Sscanf(s.Text(), "%s", &boxID)
		if err != nil {
			log.Fatal(err)
		}

		checkedLetters := make(map[rune]int)
		for _, c := range boxID {
			checkedLetters[c] = checkedLetters[c] + 1
		}

		if ruleOne(checkedLetters) {
			rule1++
		}
		if ruleTwo(checkedLetters) {
			rule2++
		}
	}

	fmt.Printf("Your checksum is: %d", (rule1 * rule2))
}

/**
* Checks wether or not the ID of the box is applicaple for the first rule
* which states that the ID must contain at least one letter exactly two times
 */
func ruleOne(checkedLetters map[rune]int) bool {
	for _, count := range checkedLetters {
		if count == 2 {
			return true
		}
	}
	return false
}

/**
* Checks wether or not the ID of the box is applicaple for the second rule
* which states that the ID must contain at least one letter exactly two times
 */
func ruleTwo(checkedLetters map[rune]int) bool {
	for _, count := range checkedLetters {
		if count == 3 {
			return true
		}
	}
	return false
}
