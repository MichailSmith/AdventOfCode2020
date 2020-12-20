package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"adventofcode.com/misha/shared"
)

func main() {
	lines := shared.ReadFileLines("day13/input.txt")

	startTime, _ := strconv.Atoi(lines[0])

	busIds := strings.Split(lines[1], ",")

	busWaitTime := make(map[int]int)
	delta := 1
	scheduleToTarget := make(map[int]int)
	firstID := -1
	for _, busIDString := range busIds {
		busInt, err := strconv.Atoi(busIDString)
		if err == nil {
			busWaitTime[busInt] = busInt - startTime%busInt

			if busWaitTime[busInt] == busInt {
				busWaitTime[busInt] = 0
			}
			if firstID < 0 {
				firstID = busInt
				delta = 0
				scheduleToTarget[busInt] = 0
			} else {
				scheduleToTarget[busInt] = busInt - delta%busInt

			}
		}
		delta++
	}

	minID := -1

	for id, waitTime := range busWaitTime {

		if minID < 0 {
			minID = id
		}

		if waitTime < busWaitTime[minID] {
			minID = id
		}
	}

	fmt.Println(busWaitTime[minID] * minID)

	vals := []int{}
	for val := range scheduleToTarget {
		vals = append(vals, val)
	}

	sort.Ints(vals)

	r := vals[len(vals)-1]
	result := scheduleToTarget[r]
	for i := len(vals) - 2; i >= 0; i-- {
		b := vals[i]
		for {
			if result%b == scheduleToTarget[b] {
				break
			}
			result += r
		}
		r *= b
	}
	fmt.Println(result)
}
