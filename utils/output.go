package utils

import (
	"fmt"
)

const (
	Debug = "DEBUG"
	Info  = "INFO"
)

// PrintWelcomeMessage prints a nice message to start you off
func PrintWelcomeMessage() {
	fmt.Print("\n\n")
	fmt.Println("=====  🎉 Welcome welcome to Advent of Code 🎉  ======")
	fmt.Println("===     Little program created by Patrick Lehmann     ==")
	fmt.Println("== You can find him at https://github.com/L33m4n123/  ==")
	fmt.Println("===                     © 2020                       ===")
	fmt.Println("=====  🎉 Welcome welcome to Advent of Code 🎉  ======")
	fmt.Print("\n\n")
}

// PrintError prints a nice Error Message to the console
func PrintError(errorMessage string) {
	fmt.Print("\n\n")
	fmt.Println(PadLeft("", len(errorMessage)/2, "⚠  "))
	fmt.Println(errorMessage)
	fmt.Println(PadLeft("", len(errorMessage)/2, "⚠  "))
	fmt.Print("\n\n")
}

// PostSolution is a wrapper to post the result of the
func PostSolution(day int, part int, answer ...interface{}) {
	fmt.Printf("========== DAY %d -- PART %d ==========\n", day, part)
	fmt.Printf("          The answer is: %v\n", answer)
	fmt.Printf("=====================================\n\n")
}

// PrintDebug prints a debug message. Used alot for me during testing x)
func PrintDebug(message string, values ...interface{}) {
	printMessage(Debug, message, values)
}

// GetPanicMessage returns a formated panic message
func GetPanicMessage(message string) string {
	return fmt.Sprintf(" DON'T PANIC \n %v", message)
}

// PrintInfo prints info Messages to the Console
func PrintInfo(message string, values ...interface{}) {
	printMessage(Info, message, values)
}

func printMessage(logLevel string, message string, values ...interface{}) {
	formattedMessage := fmt.Sprintf(message, values...)
	padLength := len(formattedMessage)
	if padLength > 50 {
		padLength = 50
	}
	entryMessage := fmt.Sprintf(" %s ", logLevel)
	entryPadLength := (padLength - len(entryMessage)) / 2
	fmt.Printf("%s %s %s\n", PadLeft("", entryPadLength, "="), entryMessage, PadRight("", entryPadLength, "="))
	fmt.Print(formattedMessage)
	fmt.Println("")
	fmt.Printf(PadLeft("\n\n", padLength+4, "="))
}
