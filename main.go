package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/l33m4n123/adventOfCodeGo/puzzlesolver"
	"github.com/l33m4n123/adventOfCodeGo/utils"
)

func main() {
	utils.PrintWelcomeMessage()
	puzzleSolver := setupPuzzleSolver()

	puzzleSolver.Run()
}

func setupPuzzleSolver() puzzlesolver.PuzzleSolver {
	var day, year int
	flags := parseFlags()

	day, year = setupDayAndYear(flags)

	puzzleSolver := puzzlesolver.New(day, year, flags)

	return puzzleSolver
}

func parseFlags() *utils.Flags {
	today := flag.Bool("today", false, "Running the test of today?")
	yesterday := flag.Bool("yesterday", false, "Running the test of today?")
	test := flag.Bool("test", false, "Running the test of today?")
	debug := flag.Bool("debug", false, "Printing debug messages [force true on test]?")
	flag.Parse()

	flags := utils.Flags{}
	flags.Test = *test
	flags.Today = *today
	flags.Yesterday = *yesterday
	flags.Debug = *debug || *test // force the debug messages to be true if I force run the tests

	if flags.Yesterday && flags.Today {
		utils.PrintInfo("You can only run either today or yesterdays puzzle. %s", "Choosing yesterdays puzzle now.")
		flags.Today = false
	}

	return &flags
}

func setupDayAndYear(flags *utils.Flags) (day, year int) {
	if flags.Today || flags.Yesterday {
		currentTime := time.Now()
		month := currentTime.Month()
		if month != time.December {
			panic(utils.GetPanicMessage("You can only use this flag during december"))
		}
		day, year = currentTime.Day(), currentTime.Year()
		if flags.Yesterday {
			day = currentTime.AddDate(0, 0, -1).Day()
		}
	} else {
		year = utils.InputNumber("Please enter the year you want to run: ")
		day = utils.InputNumber("Please enter the day you want to run: ")
	}

	if year < 2015 || year > time.Now().Year() {
		panic(utils.GetPanicMessage(fmt.Sprintf("The year must be between (including) 2015 and %d", time.Now().Year())))
	}
	if day < 1 || day > 25 {
		panic(utils.GetPanicMessage("The Advent of code only runs from the 1st to 25th of december."))
	}

	return day, year
}
