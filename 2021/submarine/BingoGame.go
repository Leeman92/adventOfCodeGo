package submarine

import (
	"fmt"
	"strconv"
	"strings"
)

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
