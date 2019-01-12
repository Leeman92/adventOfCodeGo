package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var fabric fabric
	s := bufio.NewScanner(file)
	for s.Scan() {
		var id, x, y, w, h int
		_, err := fmt.Sscanf(s.Text(), "#%d @ %d,%d: %dx%d", &id, &x, &y, &w, &h)
		if err != nil {
			log.Fatal(err)
		}
		fabric.addClaim(id, x, y, w, h)
	}

	fmt.Printf("There are %d bad claims\n", fabric.getBadClaimCount())
}

type coordinates struct {
	x, y int
}
type fabric struct {
	claims map[coordinates]int
}

func (f *fabric) addClaim(id int, x, y, w, h int) {
	if f.claims == nil {
		f.claims = make(map[coordinates]int)
	}

	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			f.claims[coordinates{x + i, y + j}]++
		}
	}
}

func (f *fabric) getBadClaimCount() int {
	count := 0
	for _, claimCount := range f.claims {
		if claimCount > 1 {
			count++
		}
	}
	return count
}
