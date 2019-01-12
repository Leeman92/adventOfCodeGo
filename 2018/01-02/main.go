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

	frequency := 0
	lookingForDuplicate := true
	var foundFrequencies []int
	for lookingForDuplicate {
		s := bufio.NewScanner(f)
		for s.Scan() {
			var n int
			_, err := fmt.Sscanf(s.Text(), "%d", &n)
			if err != nil {
				log.Fatalf("could not read %s: %v", s.Text(), err)
			}
			frequency += n
			if stringInSlice(frequency, foundFrequencies) {
				lookingForDuplicate = false
				break
			}
			foundFrequencies = append(foundFrequencies, frequency)
		}
		f.Seek(0, 0)
		if err := s.Err(); err != nil {
			log.Fatal(err)
		}
	}

	fmt.Printf("Your frequency is %d", frequency)
}

func stringInSlice(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
