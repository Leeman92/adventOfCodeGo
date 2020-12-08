package day8

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/l33m4n123/adventOfCodeGo/2020/utils"
)

// Handheld is a struct for the handhelds
type Handheld struct {
	register int
}

func (h *Handheld) changeInputToBreakLoop(allInstructions []string) {
	alreadyChanged := make(map[int]bool, len(allInstructions))
	solved := false
	counter := 0
	for !solved {
		counter++
		changedThisRun := false
		h.register = 0
		allreadyExecuted := make(map[int]bool, len(allInstructions))
		for i := 0; i < len(allInstructions); i++ {
			if allreadyExecuted[i] {
				solved = false
				break
			}

			allreadyExecuted[i] = true
			instruction := allInstructions[i]
			instructionInfo := strings.Split(instruction, " ")

			if len(instructionInfo) != 2 {
				panic(fmt.Sprintf("Could not parse instruction %v", instruction))
			}

			instrVal, convErr := strconv.Atoi(instructionInfo[1])

			if convErr != nil {
				panic(fmt.Sprintf("Could not convert %v to an integer", instructionInfo[1]))
			}

			if !changedThisRun && !alreadyChanged[i] {
				if instructionInfo[0] == "nop" && instrVal != 0 {
					instructionInfo[0] = "jmp"
					changedThisRun = true
					alreadyChanged[i] = true
				}
				if instructionInfo[0] == "jmp" {
					instructionInfo[0] = "nop"
					changedThisRun = true
					alreadyChanged[i] = true
				}
			}

			switch instructionInfo[0] {
			case "acc":
				h.register += instrVal
			case "jmp":
				i += instrVal - 1 // We run in off by ones otherwise thanks to the fact that the for loop always increases by one
			case "nop":
				fallthrough
			default:
				continue
			}

			solved = true
		}
	}

	fmt.Printf("Took %v iterations\n", counter)
}

func (h *Handheld) lookForDuplicateInstrusctions(allInstructions []string) {
	allreadyExecuted := make(map[int]bool, len(allInstructions))

	for i := 0; i < len(allInstructions); i++ {
		if allreadyExecuted[i] {
			break
		}

		allreadyExecuted[i] = true
		instruction := allInstructions[i]
		instructionInfo := strings.Split(instruction, " ")

		if len(instructionInfo) != 2 {
			panic(fmt.Sprintf("Could not parse instruction %v", instruction))
		}

		instrVal, convErr := strconv.Atoi(instructionInfo[1])

		if convErr != nil {
			panic(fmt.Sprintf("Could not convert %v to an integer", instructionInfo[1]))
		}

		switch instructionInfo[0] {
		case "acc":
			h.register += instrVal
		case "jmp":
			i += instrVal - 1 // We run in off by ones otherwise thanks to the fact that the for loop always increases by one
		case "nop":
			fallthrough
		default:
			continue
		}
	}
}

// Solve solves the puzzle
func Solve(input []string) {
	gameBoy := Handheld{}

	gameBoy.lookForDuplicateInstrusctions(input)
	utils.PostSolution(8, 1, gameBoy.register)

	gameBoy = Handheld{}
	gameBoy.changeInputToBreakLoop(input)
	utils.PostSolution(8, 2, gameBoy.register)
}
