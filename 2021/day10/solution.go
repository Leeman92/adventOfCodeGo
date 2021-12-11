package day10

import (
	"github.com/l33m4n123/adventOfCodeGo/2021/utils"
	"sort"
	"strings"
)

func Solve(lines []string) {
	syntaxScore := 0
	var incompleteLines []string
	for _, line := range lines {
		cleanup := cleanupLine(line)
		val := strings.Index(cleanup, ")")
		if val == -1 {
			val = strings.Index(cleanup, "}")
			if val == -1 {
				val = strings.Index(cleanup, "]")
				if val == -1 {
					val = strings.Index(cleanup, ">")
				}
			}
		}
		if val == -1 {
			incompleteLines = append(incompleteLines, cleanup)
			continue
		}
		syntaxScore += calculateSyntaxScoreForLine(cleanup)
	}

	utils.PostSolution(10, 1, syntaxScore)

	var scores []int
	for _, line := range incompleteLines {
		line = utils.Reverse(line)
		line = strings.ReplaceAll(line, "<", ">")
		line = strings.ReplaceAll(line, "(", ")")
		line = strings.ReplaceAll(line, "[", "]")
		line = strings.ReplaceAll(line, "{", "}")
		score := calculateCompletionScore(line)
		scores = append(scores, score)
	}

	sort.Ints(scores)
	utils.PostSolution(10, 2, scores[len(scores)/2])
}

func calculateCompletionScore(line string) int {
	score := 0
	splitLine := strings.Split(line, "")
	for _, char := range splitLine {
		score *= 5
		score += getCharValue(char)
	}
	return score
}

func getCharValue(char string) (score int) {
	switch char {
	case ")":
		score = 1
	case "]":
		score = 2
	case "}":
		score = 3
	case ">":
		score = 4
	}
	return score
}

func calculateSyntaxScoreForLine(cleanup string) int {
	bracketIndex := strings.Index(cleanup, ")")
	curlyBracketIndex := strings.Index(cleanup, "}")
	squareBracketIndex := strings.Index(cleanup, "]")
	greaterThanIndex := strings.Index(cleanup, ">")

	score := 0
	if bracketIndex == -1 {
		bracketIndex = len(cleanup) + 1
	}
	if curlyBracketIndex == -1 {
		curlyBracketIndex = len(cleanup) + 1
	}
	if squareBracketIndex == -1 {
		squareBracketIndex = len(cleanup) + 1
	}
	if greaterThanIndex == -1 {
		greaterThanIndex = len(cleanup) + 1
	}
	if bracketIndex < curlyBracketIndex && bracketIndex < squareBracketIndex && bracketIndex < greaterThanIndex {
		score = 3
	} else if curlyBracketIndex < bracketIndex && curlyBracketIndex < squareBracketIndex && curlyBracketIndex < greaterThanIndex {
		score = 1197
	} else if bracketIndex > squareBracketIndex && curlyBracketIndex > squareBracketIndex && squareBracketIndex < greaterThanIndex {
		score = 57
	} else {
		score = 25137
	}

	return score
}

func cleanupLine(line string) string {
	clean := false
	for !clean {
		length := len(line)
		line = strings.ReplaceAll(line, "<>", "")
		line = strings.ReplaceAll(line, "()", "")
		line = strings.ReplaceAll(line, "[]", "")
		line = strings.ReplaceAll(line, "{}", "")
		newLength := len(line)
		clean = length == newLength
	}
	return line
}
