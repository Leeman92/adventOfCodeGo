package submarine

import (
	"fmt"
	"github.com/l33m4n123/adventOfCodeGo/2021/utils"
	"github.com/wcharczuk/go-chart"
	"os"
	"strconv"
	"strings"
)

type Submarine struct {
	SimplePosition         utils.Coordinates
	ComplexPosition        utils.Coordinates
	SimplePositionHistory  []utils.Coordinates
	ComplexPositionHistory []utils.Coordinates
	Aim                    int
	AimHistory             []int
	InstructionSet         []Instruction
	GammaRate              []string
	EpsilonRate            []string
	OxygenRatingBinary     []string
	ScrubberRatingBinary   []string
	OxygenRating           int
	ScrubberRating         int
	LifeSupportRating      int
	PowerDraw              int
	BingoGame              BingoGame
}

type Instruction struct {
	InstructionSet func(int)
	Parameter      int
}

func (s *Submarine) PrepareNavigationComputer(lines []string, day int) {
	definedInstructions := map[int]map[string]func(int){
		2: {
			"forward": s.forward,
			"up":      s.up,
			"down":    s.down,
		},
	}

	for _, val := range lines {
		splitCommand := strings.Split(val, " ")
		if len(splitCommand) != 2 {
			panic(fmt.Sprintf("There was something off with your input. Given Input in question \"%s\"", val))
		}
		parameter, err := strconv.Atoi(splitCommand[1])
		if err != nil {
			panic(fmt.Sprintf("\"%s\" could not be parsed to an integer\nFull Instruction was \"%s\"", splitCommand[1], val))
		}

		instructionLogic, ok := definedInstructions[day][splitCommand[0]]
		if !ok {
			errorMessage := fmt.Sprintf("There was an instruction passed that has yet to be defined.\n"+
				"Passed Instruction \"%s\" is unknown for day \"%d\"", splitCommand[0], day)
			panic(errorMessage)
		}

		instr := Instruction{instructionLogic, parameter}
		s.InstructionSet = append(s.InstructionSet, instr)
	}
}

func (s *Submarine) Steer() {
	// To have the initial settings in the History with 0, 0
	s.updateHistory()
	for _, val := range s.InstructionSet {
		val.InstructionSet(val.Parameter)

		s.updateHistory()
	}
}

func (s *Submarine) GetPosition(complex bool) int {
	if !complex {
		return s.SimplePosition.X * s.SimplePosition.Y
	} else {
		return s.ComplexPosition.X * s.ComplexPosition.Y
	}
}

func (s *Submarine) RunDiagnostic(lines []string) {
	gammaRateSlice := s.calculateGammaRate(lines)
	epsilonRateSlice := s.calculateEpsilonRate(lines)
	oxygenRating, scrubberRating := s.calculateRatings(lines)

	s.GammaRate = gammaRateSlice
	s.EpsilonRate = epsilonRateSlice
	s.OxygenRatingBinary = oxygenRating
	s.OxygenRating = utils.BinarySliceToInt(oxygenRating)
	s.ScrubberRatingBinary = scrubberRating
	s.ScrubberRating = utils.BinarySliceToInt(scrubberRating)

	s.PowerDraw = utils.BinarySliceToInt(gammaRateSlice) * utils.BinarySliceToInt(epsilonRateSlice)
	s.LifeSupportRating = s.ScrubberRating * s.OxygenRating
}

func (s *Submarine) calculateGammaRate(lines []string) []string {
	bitLength := len(lines[0])
	var gammaRateSlice []string
	for i := 0; i < bitLength; i++ {
		gammaRateSlice = append(gammaRateSlice, utils.GetRelevantBit(lines, i, false))
	}

	return gammaRateSlice
}

func (s *Submarine) calculateEpsilonRate(lines []string) []string {
	bitLength := len(lines[0])
	var epsilonBinarySlice []string
	for i := 0; i < bitLength; i++ {
		epsilonBinarySlice = append(epsilonBinarySlice, utils.GetRelevantBit(lines, i, true))
	}

	return epsilonBinarySlice
}

func (s *Submarine) calculateRatings(lines []string) ([]string, []string) {
	var oxygenRating []string
	var scrubberRating []string

	oxygenReadout := lines
	scrubberReadout := lines
	for i := 0; i < len(lines[0]); i++ {
		if len(oxygenReadout) > 1 {
			mostCommonBit := utils.GetRelevantBit(oxygenReadout, i, false)
			oxygenReadout = utils.FilterSlice(oxygenReadout, mostCommonBit, i)
		}

		if len(scrubberReadout) > 1 {
			leastCommonBit := utils.GetRelevantBit(scrubberReadout, i, true)
			scrubberReadout = utils.FilterSlice(scrubberReadout, leastCommonBit, i)
		}

		if len(scrubberReadout) == 1 {
			scrubberRating = scrubberReadout
		}
		if len(oxygenReadout) == 1 {
			oxygenRating = oxygenReadout
		}

		if len(scrubberReadout) == 1 && len(oxygenReadout) == 1 {
			break
		}
	}

	return oxygenRating, scrubberRating
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

func (s *Submarine) updateHistory() {
	s.SimplePositionHistory = append(s.SimplePositionHistory, s.SimplePosition)
	s.ComplexPositionHistory = append(s.ComplexPositionHistory, s.ComplexPosition)
	s.AimHistory = append(s.AimHistory, s.Aim)
}

func (s *Submarine) DrawDepth() {
	var simpleXValues []float64
	var simpleYValues []float64
	var complexXValues []float64
	var complexYValues []float64
	for i := 0; i < len(s.SimplePositionHistory); i++ {
		simpleXValues = append(simpleXValues, float64(s.SimplePositionHistory[i].X))
		simpleYValues = append(simpleYValues, float64(s.SimplePositionHistory[i].Y*-1))
	}

	for i := 0; i < len(s.ComplexPositionHistory); i++ {
		complexXValues = append(complexXValues, float64(s.ComplexPositionHistory[i].X))
		complexYValues = append(complexYValues, float64(s.ComplexPositionHistory[i].Y*-1/100))
	}

	graph := chart.Chart{
		Series: []chart.Series{
			chart.ContinuousSeries{
				XValues: simpleXValues,
				YValues: simpleYValues,
			},
			chart.ContinuousSeries{
				XValues: complexXValues,
				YValues: complexYValues,
			},
		},
	}

	pngFile, err := os.Create("output.png")
	if err != nil {
		panic(err)
	}
	err = graph.Render(chart.PNG, pngFile)

	if err != nil {
		panic(err)
	}

	if err = pngFile.Close(); err != nil {
		panic(err)
	}
}

func (s *Submarine) GenerateBingoGame(lines []string) {
	bingoGame := BingoGame{}
	var bingoBoard BingoBoard
	for key, val := range lines {
		if key == 0 {
			generateDrawnNumbers(val, &bingoGame)
			continue
		}
		if val == "" {
			if key > 1 {
				bingoGame.BingoBoards = append(bingoGame.BingoBoards, bingoBoard)
			}
			bingoBoard = BingoBoard{}
			bingoBoard.BingoNumbers = make(map[int]map[int]BingoNumber)
			continue
		}
		bingoBoard.addLine(val)
	}

	bingoGame.BingoBoards = append(bingoGame.BingoBoards, bingoBoard)

	s.BingoGame = bingoGame

}
