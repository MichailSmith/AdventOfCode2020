package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("day1/input.txt")

	if err != nil {
		fmt.Println("Error opening file", err)
	}

	numberStrings := strings.Split(string(data), "\n")

	numbers := make([]int, len(numberStrings))

	for i := range numberStrings {
		val, err := strconv.Atoi(numberStrings[i])
		if err != nil {
			panic(err)
		}
		numbers[i] = val
	}

	length := len(numbers)

	for i := 0; i < length-1; i++ {
		for j := i + 1; j < length; j++ {
			if numbers[i]+numbers[j] == 2020 {
				fmt.Println(numbers[i] * numbers[j])
			}
		}
	}
	for i := 0; i < length-2; i++ {
		for j := i + 1; j < length-1; j++ {
			for k := j + 1; k < length; k++ {
				if numbers[i]+numbers[j]+numbers[k] == 2020 {
					fmt.Println(numbers[i] * numbers[j] * numbers[k])
				}
			}
		}
	}
}
