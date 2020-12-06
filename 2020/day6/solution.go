package day6

import (
	"strings"

	"github.com/l33m4n123/adventOfCodeGo/2020/utils"
)

// Solve runs the puzzle
func Solve(lines []string) {
	allAnswers := parseGroups(lines)
	formattedAnswer := formatGroupsForPartOne(allAnswers)
	uniqueGroupAnswers := utils.UniqueMultipleStringSlice(formattedAnswer)

	partOne(uniqueGroupAnswers)
	partTwo(allAnswers)
}

func partOne(answers [][]string) {
	totalAnswers := 0
	for _, answer := range answers {
		totalAnswers += len(answer)
	}

	utils.PostSolution(6, 1, totalAnswers)
}

func partTwo(answers [][][]string) {
	allAnswersYes := 0
	for _, groups := range answers {
		expectedAnswerCount := len(groups)
		answers := make(map[string]int)
		for _, personAnswer := range groups {
			for _, answer := range personAnswer {
				answers[answer]++
			}
		}

		for _, value := range answers {
			if value != expectedAnswerCount {
				continue
			}
			allAnswersYes++
		}
	}

	utils.PostSolution(6, 2, allAnswersYes)
}

func formatGroupsForPartOne(allAnswers [][][]string) [][]string {
	formatedAnswers := [][]string{}
	for _, groupAnswers := range allAnswers {
		groupAnswer := []string{}
		for _, answer := range groupAnswers {
			for _, value := range answer {
				groupAnswer = append(groupAnswer, value)
			}
		}
		formatedAnswers = append(formatedAnswers, groupAnswer)
	}

	return formatedAnswers
}

func parseGroups(lines []string) [][][]string {
	groupAnswer := [][]string{}
	answersByGroup := [][][]string{}

	for _, answers := range lines {
		if answers == "" {
			answersByGroup = append(answersByGroup, groupAnswer)
			groupAnswer = [][]string{}
			continue
		}
		answerSplit := strings.Split(answers, "")
		groupAnswer = append(groupAnswer, answerSplit)
	}

	answersByGroup = append(answersByGroup, groupAnswer)

	return answersByGroup
}
