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
	SimplePosition         Coordinates
	ComplexPosition        Coordinates
	SimplePositionHistory  []Coordinates
	ComplexPositionHistory []Coordinates
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

type Coordinates struct {
	X int
	Y int
}

type Instruction struct {
	InstructionSet func(int)
	Parameter      int
}

type BingoGame struct {
	NumbersDrawn       []int
	BingoBoards        []BingoBoard
	currentRound       int
	Done               bool
	FirstWinningBoard  BingoBoard
	FirstWinningNumber int
	LastWinningBoard   BingoBoard
	LastWinningNumber  int
	FirstWin           bool
}

type BingoBoard struct {
	BingoNumbers map[int]map[int]BingoNumber
	currentLine  int
	Done         bool
}

type BingoNumber struct {
	Position Coordinates
	Value    int
	Marked   bool
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

func (board *BingoBoard) addLine(line string) {
	numberStringSlice := strings.Split(line, " ")
	counter := 0
	board.BingoNumbers[board.currentLine] = make(map[int]BingoNumber)
	for _, val := range numberStringSlice {
		if strings.Trim(val, " ") == "" {
			continue
		}

		num, err := strconv.Atoi(strings.Trim(val, " "))
		if err != nil {
			panic(err)
		}

		coordinate := Coordinates{counter, board.currentLine}
		bingoNumber := BingoNumber{coordinate, num, false}
		board.BingoNumbers[board.currentLine][counter] = bingoNumber

		counter++
	}

	board.currentLine += 1
}

func generateDrawnNumbers(val string, bingoGame *BingoGame) {
	numbersDrawnStringSlice := strings.Split(val, ",")
	for _, stringNumber := range numbersDrawnStringSlice {
		num, err := strconv.Atoi(stringNumber)
		if err != nil {
			panic(err)
		}
		bingoGame.NumbersDrawn = append(bingoGame.NumbersDrawn, num)
	}
}

func (game *BingoGame) Play() {
	game.FirstWin = false
	for !game.Done {
		game.nextRound()
	}
}

func (game *BingoGame) nextRound() {
	currentRound := game.currentRound
	if currentRound >= len(game.NumbersDrawn) {
		fmt.Println("No more rounds to play...")
		game.Done = true
		return
	}

	currentNumber := game.NumbersDrawn[currentRound]
	for key, board := range game.BingoBoards {
		if board.Done {
			continue
		}
		board.MarkNumber(currentNumber)
		game.BingoBoards[key] = board
		if !board.CheckCompleteness() {
			continue
		}
		board.Done = true
		game.BingoBoards[key] = board

		if !game.FirstWin {
			game.FirstWinningBoard = board
			game.FirstWinningNumber = currentNumber
			game.FirstWin = true
		}
		game.LastWinningBoard = board
		game.LastWinningNumber = currentNumber
	}

	game.currentRound += 1
}

func (board *BingoBoard) MarkNumber(number int) {
	for x := 0; x < 5; x++ {
		for y := 0; y < 5; y++ {
			bingoEntry := board.BingoNumbers[y][x]
			if bingoEntry.Value != number {
				continue
			}
			bingoEntry.Marked = true
			board.BingoNumbers[y][x] = bingoEntry
		}
	}
}

func (board *BingoBoard) CheckCompleteness() bool {
	rows := board.checkRows()
	columns := board.checkColumns()
	return rows || columns
}

func (board *BingoBoard) checkRows() (completed bool) {
	completionCount := 0
	for y := 0; y < 5; y++ {
		if completed && completionCount == 5 {
			return completed
		}
		completionCount = 0
		completed = false
		for x := 0; x < 5; x++ {
			entry := board.BingoNumbers[y][x]
			if !entry.Marked {
				completed = false
				break
			}
			completionCount += 1
			completed = true
		}
	}
	return completed
}

func (board *BingoBoard) checkColumns() (completed bool) {
	completionCount := 0
	for y := 0; y < 5; y++ {
		if completed && completionCount == 5 {
			return completed
		}
		completionCount = 0
		completed = false
		for x := 0; x < 5; x++ {
			entry := board.BingoNumbers[x][y]
			if !entry.Marked {
				completed = false
				break
			}
			completionCount += 1
			completed = true
		}
	}
	return completed
}

func (board *BingoBoard) CalculateMarkedSum() int {
	return board.calculateSum(true)
}

func (board *BingoBoard) CalculateUnmarkedSum() int {
	return board.calculateSum(false)
}

func (board *BingoBoard) calculateSum(marked bool) int {
	sum := 0
	for x := 0; x < 5; x++ {
		for y := 0; y < 5; y++ {
			entry := board.BingoNumbers[x][y]
			if entry.Marked == marked {
				sum += entry.Value
			}
		}
	}

	return sum
}

func (board *BingoBoard) PrintDebug() {
	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			spacerLeft := "| "
			spacerRight := " |"
			if board.BingoNumbers[y][x].Marked {
				spacerLeft = "|>"
				spacerRight = "<|"
			}
			fmt.Printf("%s%02d%s", spacerLeft, board.BingoNumbers[y][x].Value, spacerRight)
		}
		fmt.Printf("\n")
	}
}
