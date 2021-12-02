package submarine

import (
	"strconv"
	"strings"
)

type Submarine struct {
	SimplePosition Coordinates
	ComplexPosition Coordinates
	Aim      int
	InstructionSet []string
}

type Coordinates struct {
	X int
	Y int
}

func (s *Submarine) Steer() {
	for _, val := range s.InstructionSet {
		splitString := strings.Split(val, " ")
		if len(splitString) != 2 {
			continue
		}
		count, _ := strconv.Atoi(splitString[1])
		if strings.Index(val, "forward") == 0 {
			s.SimplePosition.X += count
			s.ComplexPosition.X += count
			s.ComplexPosition.Y += count * s.Aim
		}
		if strings.Index(val, "up") == 0 {
			s.SimplePosition.Y -= count
			s.Aim -= count
		}
		if strings.Index(val, "down") == 0 {
			s.SimplePosition.Y += count
			s.Aim += count
		}
	}
}

func (s *Submarine) GetPosition(complex bool) int {
	if !complex {
		return s.SimplePosition.X * s.SimplePosition.Y
	} else {
		return s.ComplexPosition.X * s.ComplexPosition.Y
	}
}