package main

import "fmt"

func main() {
	input := []int{9, 12, 1, 4, 17, 0, 18}
	state := make(map[int]int)

	currentNumber := 0
	previousNumber := 0
	for i := 0; i < 2020; i++ {
		lastAppearance, exists := state[previousNumber]
		if exists {
			currentNumber = i - lastAppearance - 1
		}
		if !exists {
			currentNumber = 0
		}
		if i < len(input) {
			currentNumber = input[i]
		}
		state[previousNumber] = i - 1
		previousNumber = currentNumber
	}
	fmt.Println(currentNumber)

	for i := 2020; i < 30000000; i++ {
		lastAppearance, exists := state[previousNumber]
		if exists {
			currentNumber = i - lastAppearance - 1
		}
		if !exists {
			currentNumber = 0
		}
		if i < len(input) {
			currentNumber = input[i]
		}
		state[previousNumber] = i - 1
		previousNumber = currentNumber
	}
	fmt.Println(currentNumber)
}
