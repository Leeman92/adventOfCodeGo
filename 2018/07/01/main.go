package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

var instructions map[string][]string

func main() {
	fmt.Println("=========== START ==========")
	instructions = make(map[string][]string)
	b, err := ioutil.ReadFile("testinput.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(b), "\n")

	parseLines(lines)
}

func parseLines(lines []string) {
	for _, line := range lines {
		var firstInstruction, secondInstruction string
		fmt.Sscanf(line, "Step %s must be finished before step %s can begin.", &firstInstruction, &secondInstruction)
		fmt.Printf("%s -> %s\n", firstInstruction, secondInstruction)

		if _, ok := instructions[firstInstruction]; !ok {
			instructions[firstInstruction] = append([]string{""}, secondInstruction)
		} else {
			instructions[firstInstruction] = append(instructions[firstInstruction], secondInstruction)
		}
	}

	fmt.Printf("%v", instructions)
}
