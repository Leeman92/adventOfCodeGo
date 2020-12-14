package utils

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

// InputString Takes a question that is used to request a string from stdin
func InputString(question string) string {
	fmt.Print("❔ ", question, ": ")
	reader := bufio.NewReader(os.Stdin)
	responseBytes, _, _ := reader.ReadLine()

	return string(responseBytes)
}

// InputNumber Takes a question that is used to request an int from stdin
func InputNumber(question string) int {
	var responseNumber int
	err := errors.New("no input")

	for err != nil {
		PrintError(err.Error())
		response := InputString(question)

		responseNumber, err = strconv.Atoi(response)
	}

	return responseNumber
}
