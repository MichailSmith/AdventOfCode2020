package main

import (
	"fmt"

	"adventofcode.com/misha/shared"
)

func main() {
	lines := shared.ReadFileBytes("day11/input.txt")

	firstSet := lines
	for {
		next, changed := getNextIteration(firstSet, getNewValue)
		firstSet = next
		if !changed {
			break
		}
	}

	fmt.Println(countOccupied(firstSet))

	secondSet := lines
	for {
		next, changed := getNextIteration(secondSet, getNewValueLookFar)
		secondSet = next
		if !changed {
			break
		}
	}
	fmt.Println(countOccupied(secondSet))
}

func getNewValueLookFar(y int, x int, grid [][]byte) (byte, bool) {
	directions := []func(int, int) (int, int){
		func(i int, j int) (int, int) {
			return i + 1, j
		},
		func(i int, j int) (int, int) {
			return i + 1, j + 1
		},
		func(i int, j int) (int, int) {
			return i + 1, j - 1
		},
		func(i int, j int) (int, int) {
			return i - 1, j + 1
		},
		func(i int, j int) (int, int) {
			return i - 1, j
		},
		func(i int, j int) (int, int) {
			return i - 1, j - 1
		},
		func(i int, j int) (int, int) {
			return i, j + 1
		},
		func(i int, j int) (int, int) {
			return i, j - 1
		},
	}

	if grid[y][x] == '.' {
		return '.', false
	}
	count := 0
	for _, iterator := range directions {
		nextY, nextX := iterator(y, x)
		for {
			if nextX < 0 ||
				nextY < 0 ||
				nextY >= len(grid) ||
				nextX >= len(grid[nextY]) ||
				grid[nextY][nextX] == 'L' {
				break
			}
			if grid[nextY][nextX] == '#' {
				count++
				break
			}
			nextY, nextX = iterator(nextY, nextX)
		}
	}
	if count == 0 {
		return '#', grid[y][x] != '#'
	}
	if count >= 5 {
		return 'L', grid[y][x] != 'L'
	}

	return grid[y][x], false
}

func getNewValue(y int, x int, grid [][]byte) (byte, bool) {
	if grid[y][x] == '.' {
		return '.', false
	}
	adjacentCount := 0
	for i := shared.Max(y-1, 0); i < shared.Min(y+2, len(grid)); i++ {
		for j := shared.Max(x-1, 0); j < shared.Min(x+2, len(grid[y])); j++ {
			if i == y && j == x {
				continue
			}
			if grid[i][j] == '#' {
				adjacentCount++
			}
		}
	}
	if adjacentCount == 0 {
		return '#', grid[y][x] != '#'
	}
	if adjacentCount >= 4 {
		return 'L', grid[y][x] != 'L'
	}
	return grid[y][x], false
}

func getNextIteration(grid [][]byte, newValueFunc func(int, int, [][]byte) (byte, bool)) ([][]byte, bool) {
	next := make([][]byte, len(grid))
	anyChanged := false
	for i, row := range grid {
		next[i] = make([]byte, len(row))
		for j := range row {
			newValue, valueCanged := newValueFunc(i, j, grid)
			anyChanged = anyChanged || valueCanged
			next[i][j] = newValue
		}
	}
	return next, anyChanged
}

func countOccupied(grid [][]byte) int {
	count := 0
	for _, row := range grid {
		for _, value := range row {
			if value == '#' {
				count++
			}
		}
	}
	return count
}
