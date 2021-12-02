package submarine

import (
	"fmt"
	"strconv"
	"strings"
)

type Submarine struct {
	SimplePosition Coordinates
	ComplexPosition Coordinates
	Aim      int
	InstructionSet []Instruction

}

type Coordinates struct {
	X int
	Y int
}

type Instruction struct {
	Something func(int)
	Parameter int
}

func (s *Submarine) PrepareNavigationComputer(lines []string) {
	definedInstructions := map[string]func(int) {
		"forward": s.forward,
		"up": s.up,
		"down": s.down,
	}

	for _, val := range lines {
		splitCommand := strings.Split(val, " ")
		if len(splitCommand) != 2 {
			panic(fmt.Sprintf("There was something off with your input. Given Input in question %s", val))
		}
		parameter, err := strconv.Atoi(splitCommand[1])
		if err != nil {
			panic(fmt.Sprintf("%s could not be parsed to an integer", splitCommand[1]))
		}

		instructionLogic, ok := definedInstructions[splitCommand[0]]
		if !ok {
			errorMessage := fmt.Sprintf("There was an instruction passed that has yet to be defined." +
				"Passed Instruction %s is unkown", splitCommand[0])
			panic(errorMessage)
		}

		s.InstructionSet = append(s.InstructionSet, Instruction{instructionLogic, parameter})
	}
}

func (s *Submarine) Steer() {
	for _, val := range s.InstructionSet {
		val.Something(val.Parameter)
	}
}

func (s *Submarine) GetPosition(complex bool) int {
	if !complex {
		return s.SimplePosition.X * s.SimplePosition.Y
	} else {
		return s.ComplexPosition.X * s.ComplexPosition.Y
	}
}


func (s *Submarine) forward(amount int) {
	s.SimplePosition.X += amount
	s.ComplexPosition.X += amount
	s.ComplexPosition.Y += amount * s.Aim
}

func (s *Submarine) up(amount int) {
	s.SimplePosition.Y -= amount
	s.Aim -= amount
}

func (s *Submarine) down(amount int) {
	s.SimplePosition.Y += amount
	s.Aim += amount
}