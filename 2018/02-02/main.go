package main

import (
	"bufio"
	"bytes"
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

	var boxIDSlice []string
	s := bufio.NewScanner(f)
	for s.Scan() {
		var boxID string
		_, err := fmt.Sscanf(s.Text(), "%s", &boxID)
		if err != nil {
			log.Fatal(err)
		}
		boxIDSlice = append(boxIDSlice, boxID)
	}

	/*
	 * Now loop over each and compare every other boxID to have only ONE difference.
	 * if the two IDs are found that only differ by one letter, get its common letters
	 */

	IDsFound := false
	var IDOne, IDTwo string
	for _, boxID := range boxIDSlice {
		for _, otherBoxID := range boxIDSlice {
			if ruleApplies(boxID, otherBoxID) {
				IDsFound = true
				IDOne = boxID
				IDTwo = otherBoxID
				break
			}
		}
		if IDsFound {
			break
		}
	}
	answer := getCommonChars(IDOne, IDTwo)
	fmt.Printf("The common letters are: %s", answer)
}

/*
 * Checks every ID and makes sure they differ in only one letter (make sure to do not count itself as one)
 */
func ruleApplies(boxIDOne string, boxIDTwo string) bool {
	boxIDTwoSlice := []rune(boxIDTwo)
	counter := 0
	for pos, char := range boxIDOne {
		if char != boxIDTwoSlice[pos] {
			counter++
		}
	}
	if counter == 0 || counter > 1 {
		return false
	}
	return true
}

/*
 * Compares the two strings and returns all common characters
 */
func getCommonChars(IDOne string, IDTwo string) string {
	var answer bytes.Buffer
	boxIDTwoSlice := []rune(IDTwo)
	for pos, char := range IDOne {
		if char == boxIDTwoSlice[pos] {
			answer.WriteString(string(char))
		}
	}

	return answer.String()
}
