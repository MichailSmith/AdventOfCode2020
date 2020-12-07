package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"adventofcode.com/misha/shared"
)

func main() {
	ruleLines := shared.ReadFileLines("day7/input.txt")

	regex := regexp.MustCompile(`^?(.*) bags? contain (.*)\.$`)

	colorRegex := regexp.MustCompile(`^\s?(\d+)\s(.*) bags?$`)

	canBeInside := map[string]map[string]bool{}

	mustContain := map[string]map[string]int{}

	bags := make([]string, len(ruleLines))
	for i, ruleLine := range ruleLines {
		matches := regex.FindStringSubmatch(ruleLine)

		if len(matches) == 3 {
			bags[i] = matches[1]
			innerColors := strings.Split(matches[2], ",")
			mustContain[matches[1]] = map[string]int{}

			for _, colorString := range innerColors {
				values := colorRegex.FindStringSubmatch(colorString)
				if len(values) == 3 {
					val, err := strconv.Atoi(values[1])
					if err != nil {
						panic(err)
					}
					mustContain[matches[1]][values[2]] = val
					_, exists := canBeInside[values[2]]
					if !exists {
						canBeInside[values[2]] = map[string]bool{}
					}
					canBeInside[values[2]][matches[2]] = true
				}
			}
		}
	}
	containerOptions := getContainerOptions(canBeInside, map[string]bool{"shiny gold": true}, "shiny gold")

	fmt.Println(len(containerOptions) - 1)

	fmt.Println(getRequiredCount(mustContain, "shiny gold") - 1)
}

func getContainerOptions(canBeInside map[string]map[string]bool, alreadyVisited map[string]bool, target string) map[string]bool {
	options := canBeInside[target]

	newOptions := []string{}

	for value := range options {
		_, exists := alreadyVisited[value]
		if !exists {
			newOptions = append(newOptions, value)
			alreadyVisited[value] = true
		}
	}

	for _, value := range newOptions {
		children := getContainerOptions(canBeInside, alreadyVisited, value)
		for child := range children {
			alreadyVisited[child] = true
		}
	}

	return alreadyVisited
}

func getRequiredCount(mustContain map[string]map[string]int, target string) int {
	count := 1
	children := mustContain[target]

	for child, childCount := range children {
		count += getRequiredCount(mustContain, child) * childCount
	}

	return count
}
