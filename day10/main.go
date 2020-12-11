package main

import (
	"fmt"
	"sort"
	"strconv"

	"adventofcode.com/misha/shared"
)

func main() {
	lines := shared.ReadFileLines("day10/input.txt")

	jolts := make([]int, len(lines)+1)
	jolts[0] = 0
	for i := range lines {
		value, err := strconv.Atoi(lines[i])
		if err != nil {
			panic(err)
		}
		jolts[i+1] = value
	}
	sort.Ints(jolts)

	differences := map[int]int{}
	for i := 1; i < len(jolts); i++ {
		difference := jolts[i] - jolts[i-1]
		differences[difference]++
	}
	fmt.Println((differences[1]) * (differences[3] + 1))

	count := 1
	subConnections := splitConnections(jolts)
	possibilityChan := make(chan int)
	for _, subset := range subConnections {
		go func(channel chan int, connections []int) {
			channel <- findConfigurations(connections)
		}(possibilityChan, subset)
	}

	for range subConnections {
		number := <-possibilityChan
		count *= number
	}

	fmt.Println(count)
}

func splitConnections(jolts []int) [][]int {
	result := [][]int{}
	startIndex := 0
	for i := 1; i < len(jolts); i++ {
		difference := jolts[i] - jolts[i-1]
		if difference == 3 {
			result = append(result, jolts[startIndex:i])
			startIndex = i
		}
	}

	if startIndex < len(jolts)-1 {
		result = append(result, jolts[startIndex:])
	}
	return result
}

func findConfigurations(jolts []int) int {
	if len(jolts) < 3 {
		return 1
	}

	possibilities := findConfigurations(jolts[1:])

	if jolts[2]-jolts[1] <= 3 {
		possibilities += findConfigurations(jolts[2:])
	}
	if len(jolts) > 3 && jolts[3]-jolts[1] <= 3 {
		possibilities += findConfigurations(jolts[3:])
	}

	return possibilities
}
