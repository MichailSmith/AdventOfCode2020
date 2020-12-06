package main

import (
	"fmt"

	"adventofcode.com/misha/shared"
)

func main() {
	rowStrings := shared.ReadFileLines("day3/input.txt")

	rows := make([][]byte, len(rowStrings))
	for i := range rowStrings {
		rows[i] = []byte(rowStrings[i])
	}

	treeCounts := [5]int{
		countTrees(rows, 1, 1),
		countTrees(rows, 3, 1),
		countTrees(rows, 5, 1),
		countTrees(rows, 7, 1),
		countTrees(rows, 1, 2),
	}
	fmt.Println(treeCounts[1])

	total := 1
	for i := range treeCounts {
		total *= treeCounts[i]
	}

	fmt.Println(total)
}

func countTrees(rows [][]byte, xSlope int, ySlope int) int {
	treeCount := 0
	x := 0
	for y := 0; y < len(rows); y += ySlope {
		row := rows[y]
		if row[x%len(row)] == '#' {
			treeCount++
		}
		x += xSlope
	}
	return treeCount
}
