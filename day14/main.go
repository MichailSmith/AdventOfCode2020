package main

import (
	"fmt"
	"regexp"
	"strconv"

	"adventofcode.com/misha/shared"
)

func main() {
	lines := shared.ReadFileLines("day14/input.txt")
	mem := make(map[int]int)
	mem2 := make(map[int]int)
	maskRegex := regexp.MustCompile(`mask = ([10X]+)`)
	memRegex := regexp.MustCompile(`mem\[(\d+)\] = (\d+)`)
	andMask := 0
	orMask := 0
	memoryOrMasks := []int{}
	memoryAndMasks := []int{}
	for _, line := range lines {
		maskMatches := maskRegex.FindStringSubmatch(line)
		if len(maskMatches) == 2 {
			andMask, orMask = getMasks(maskMatches[1])
			memoryOrMasks, memoryAndMasks = getMemoryAddressMasks(maskMatches[1])
		}
		memMatches := memRegex.FindStringSubmatch(line)

		if len(memMatches) == 3 {
			index, _ := strconv.Atoi(memMatches[1])
			value, _ := strconv.Atoi(memMatches[2])

			mem[index] = (value & andMask) | orMask
			for i, mask := range memoryOrMasks {
				newIndex := (index & memoryAndMasks[i]) | mask
				mem2[newIndex] = value
			}
		}
	}

	total := 0
	for _, value := range mem {
		total += value
	}
	fmt.Println(total)

	total2 := 0
	for _, value := range mem2 {
		total2 += value
	}
	fmt.Println(total2)
}

func getMasks(mask string) (andMask int, orMask int) {
	andMask = 0
	orMask = 0

	for _, val := range mask {
		andMask = andMask << 1
		orMask = orMask << 1
		switch val {
		case '1':
			orMask++
			andMask++
		case 'X':
			andMask++
		}
	}
	return andMask, orMask
}

func getMemoryAddressMasks(mask string) ([]int, []int) {
	orMasks := []int{0}
	andMasks := []int{0}
	for _, value := range mask {
		switch value {
		case '1':
			for i := range orMasks {
				orMasks[i] = orMasks[i] << 1
				orMasks[i]++
				andMasks[i] = andMasks[i] << 1
				andMasks[i]++
			}
		case '0':
			for i := range orMasks {
				orMasks[i] = orMasks[i] << 1
				andMasks[i] = andMasks[i] << 1
				andMasks[i]++
			}
		case 'X':
			toAppendOr := make([]int, len(orMasks))
			toAppendAnd := make([]int, len(andMasks))
			for i := range orMasks {
				orMasks[i] = orMasks[i] << 1
				toAppendOr[i] = orMasks[i] + 1
				andMasks[i] = andMasks[i] << 1
				toAppendAnd[i] = andMasks[i] + 1
			}
			orMasks = append(orMasks, toAppendOr...)
			andMasks = append(andMasks, toAppendAnd...)

		}
	}
	return orMasks, andMasks
}
