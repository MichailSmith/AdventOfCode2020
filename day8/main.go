package main

import (
	"fmt"
	"regexp"
	"strconv"

	"adventofcode.com/misha/shared"
)

func main() {
	lines := shared.ReadFileLines("day8/input.txt")

	fmt.Println(runProgram(lines))
	fmt.Println(fixProgram(lines))
}

func runProgram(lines []string) (int, bool) {
	cmdRegex := regexp.MustCompile(`(.*) ([+-]\d+)`)
	accumulator := 0
	visited := map[int]bool{}
	i := 0
	terminated := false
	for {
		if i == len(lines) {
			terminated = true
			break
		}
		if visited[i] {
			break
		}
		visited[i] = true
		values := cmdRegex.FindStringSubmatch(lines[i])

		if len(values) != 3 {
			panic(fmt.Sprintf("%v: %v", lines[i], values))
		}

		cmd := values[1]
		literal, err := strconv.Atoi(values[2])

		if err != nil {
			panic(err)
		}

		switch cmd {
		case "jmp":
			i += literal
		case "acc":
			accumulator += literal
			fallthrough
		case "nop":
			i++
		}
	}
	return accumulator, terminated
}

func fixProgram(lines []string) int {
	cmdRegex := regexp.MustCompile(`(.*) ([+-]\d+)`)
	newLines := make([]string, len(lines))
	for i, line := range lines {
		values := cmdRegex.FindStringSubmatch(line)

		if len(values) != 3 {
			panic(fmt.Sprintf("%v: %v", line, values))
		}

		cmd := values[1]

		switch cmd {
		case "jmp":
			copy(newLines, lines)
			newLines[i] = fmt.Sprintf("nop %s", values[2])
			value, terminated := runProgram(newLines)
			if terminated {
				return value
			}
		case "nop":
			copy(newLines, lines)
			newLines[i] = fmt.Sprintf("jmp %s", values[2])
			value, terminated := runProgram(newLines)
			if terminated {
				return value
			}
		}
	}

	panic("never terminated!")
}
