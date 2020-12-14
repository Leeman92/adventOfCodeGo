package puzzlesolver

import (
	"fmt"

	"github.com/l33m4n123/adventOfCodeGo/utils"
)

// PuzzleSolver struct to keep data relevant for puzzle
type PuzzleSolver struct {
	day, year   int
	debug, test bool
}

// Run 's the puzzle for the given day
func (solver *PuzzleSolver) Run() {
	fmt.Println("Run the given puzzle. If it doesn't exist yet, create a skelleton working puzzle structure for it")
	utils.PrintDebug("%v", *solver)
}

// New returns a new PuzzleSolver struct
func New(day, year int, flags *utils.Flags) PuzzleSolver {
	return PuzzleSolver{day: day, year: year, debug: flags.Test, test: flags.Test}
}
