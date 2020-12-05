package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("day5/input.txt")

	if err != nil {
		fmt.Println("Error opening file", err)
	}

	numberStrings := strings.Split(string(data), "\n")

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
