package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	fmt.Println("=========== START ==========")
	b, err := ioutil.ReadFile("testinput.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(b), "\n")

	for _, line := range lines {
		fmt.Printf("%v\n", line)
	}
}
