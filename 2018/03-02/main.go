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
	fmt.Printf("The best claim option is claim #%d", fabric.getBestClaim())
}

type coordinates struct {
	x, y int
}
type fabric struct {
	ids    map[int]bool
	claims map[coordinates][]int
}

func (f *fabric) addClaim(id, x, y, w, h int) {
	if f.claims == nil {
		f.claims = make(map[coordinates][]int)
	}

	if f.ids == nil {
		f.ids = make(map[int]bool)
	}

	f.ids[id] = true

	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			coordinate := coordinates{x + i, y + j}
			f.claims[coordinate] = append(f.claims[coordinate], id)
		}
	}
}

func (f *fabric) getBadClaimCount() int {
	count := 0
	for _, ids := range f.claims {
		if len(ids) > 1 {
			count++
		}
	}
	return count
}

func (f *fabric) getBestClaim() int {
	possibleBestClaims := f.ids
	for _, ids := range f.claims {
		if len(ids) <= 1 {
			continue
		}
		for _, id := range ids {
			delete(possibleBestClaims, id)
		}
		if len(possibleBestClaims) <= 1 {
			break
		}
	}

	if len(possibleBestClaims) != 1 {
		log.Fatalf("We screwed up bigtimes! There should be only one claim. We have %d", len(possibleBestClaims))
	}

	for id := range possibleBestClaims {
		return id
	}

	return 0
}
