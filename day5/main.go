package main

import (
	"fmt"
	"sort"

	"adventofcode.com/misha/shared"
)

func main() {
	numberStrings := shared.ReadFileLines("day5/input.txt")

	numbers := make([]int, len(numberStrings))

	for i := range numberStrings {
		numbers[i] = 0
		for j := range numberStrings[i] {
			numbers[i] = numbers[i] << 1
			if numberStrings[i][j] == 'B' || numberStrings[i][j] == 'R' {
				numbers[i]++
			}
		}
	}

	sort.Ints(numbers)
	max := 0
	for i := range numbers {
		if numbers[i] > max {
			max = numbers[i]
		}
	}

	fmt.Println(max)
	sort.Ints(numbers)

	previous := -1

	for i := range numbers {
		if previous == numbers[i]-2 {
			break
		}
		previous = numbers[i]
	}

	fmt.Println(previous + 1)
}
