package main

import (
	"fmt"
	"strconv"

	"adventofcode.com/misha/shared"
)

func main() {
	lines := shared.ReadFileLines("day9/input.txt")

	values := make([]int, len(lines))
	for i, line := range lines {
		value, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		values[i] = value
	}

	invalid := getFirstInvalid(values)
	fmt.Println(invalid)
	fmt.Println(getWeakness(values, invalid))
}

func getFirstInvalid(values []int) int {
	window := make([]int, 25)

	for i, value := range values {
		if i >= len(window) {
			valid := false
		VALIDATOR:
			for j := range window {
				for k := j + 1; k < len(window); k++ {
					if window[k]+window[j] == value && window[k] != window[j] {
						valid = true
						continue VALIDATOR
					}
				}
			}
			if !valid {
				return value
			}
		}
		window = append(window[1:], value)
	}

	panic("all numbers valid")
}

func getWeakness(values []int, invalid int) int {
	for i := range values {
		acc := values[i]
		min := acc
		max := acc

		for j := i + 1; j < len(values); j++ {
			current := values[j]
			if min > current {
				min = current
			}
			if max < current {
				max = current
			}
			acc += current

			if acc == invalid {
				return min + max
			}
			if acc > invalid {
				continue
			}
		}
	}
	panic("no weakness found")
}
