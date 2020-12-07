package main

import (
	"fmt"

	"adventofcode.com/misha/shared"
)

func main() {
	answerLines := shared.ReadFileLines("day6/input.txt")

	answerGroups := make([][]string, len(answerLines))

	currentAnswers := []string{}
	for i := range answerLines {
		if len(answerLines[i]) == 0 {
			answerGroups = append(answerGroups, currentAnswers)
			currentAnswers = []string{}
		} else {
			currentAnswers = append(currentAnswers, answerLines[i])
		}
	}

	runningTotal := 0
	anyoneAnswered := map[byte]bool{}
	for i := range answerGroups {
		anyoneAnswered = map[byte]bool{}
		for j := range answerGroups[i] {
			for k := range answerGroups[i][j] {
				anyoneAnswered[answerGroups[i][j][k]] = true
			}
		}
		runningTotal += len(anyoneAnswered)
	}

	fmt.Println(runningTotal)

	runningTotal = 0
	everyoneAnswered := map[byte]int{}
	for i := range answerGroups {
		everyoneAnswered = map[byte]int{}
		for j := range answerGroups[i] {
			for k := range answerGroups[i][j] {
				everyoneAnswered[answerGroups[i][j][k]]++
			}
		}
		for _, value := range everyoneAnswered {
			if value == len(answerGroups[i]) {
				runningTotal++
			}
		}
	}

	fmt.Println(runningTotal)

}
