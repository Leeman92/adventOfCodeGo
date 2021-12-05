package submarine

import (
	"fmt"
	"github.com/l33m4n123/adventOfCodeGo/2021/utils"
	"strconv"
	"strings"
)

type BingoBoard struct {
	BingoNumbers map[int]map[int]BingoNumber
	currentLine  int
	Done         bool
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

		coordinate := utils.Coordinates{X: counter, Y: board.currentLine}
		bingoNumber := BingoNumber{coordinate, num, false}
		board.BingoNumbers[board.currentLine][counter] = bingoNumber

		counter++
	}

	board.currentLine += 1
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
