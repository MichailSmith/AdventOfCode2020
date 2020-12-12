package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"

	"adventofcode.com/misha/shared"
)

type instruction struct {
	action string
	amount int
}

func main() {
	lines := shared.ReadFileLines("day12/input.txt")

	fmt.Println(lines)

	regex := regexp.MustCompile(`(\w)(\d+)`)

	instructions := make([]instruction, len(lines))

	for i, value := range lines {
		matches := regex.FindStringSubmatch(value)

		amount, _ := strconv.Atoi(matches[2])
		if len(matches) == 3 {
			instructions[i] = instruction{
				action: matches[1],
				amount: amount,
			}
		}
	}

	x := 0
	y := 0
	facing := 0
	for _, ins := range instructions {
		x, y, facing = executeInstruction(x, y, facing, ins)
	}

	fmt.Println(math.Abs(float64(x)) + math.Abs(float64(y)))

	x = 0
	y = 0
	wpx := 10
	wpy := 1
	for _, ins := range instructions {
		x, y, wpx, wpy = executeInstruction2(x, y, wpx, wpy, ins)
	}

	fmt.Println(math.Abs(float64(x)) + math.Abs(float64(y)))
}

func executeInstruction(x int, y int, facing int, ins instruction) (int, int, int) {

	switch ins.action {
	case "E":
		x += ins.amount
	case "W":
		x -= ins.amount
	case "N":
		y += ins.amount
	case "S":
		y -= ins.amount
	case "F":
		switch facing {
		case 0:
			x += ins.amount
		case 90:
			y += ins.amount
		case 180:
			x -= ins.amount
		case 270:
			y -= ins.amount
		}
	case "L":
		facing = rotate(ins.amount, facing)
	case "R":
		facing = rotate(-ins.amount, facing)
	}

	return x, y, facing
}

func rotate(deg int, facing int) int {
	newF := facing + deg
	if newF < 0 {
		newF += 360
	}
	return newF % 360
}

func rotate2(deg int, wpx int, wpy int) (int, int) {

	switch deg {
	case 90:
		fallthrough
	case -270:
		return -wpy, wpx
	case 180:
		fallthrough
	case -180:
		return -wpx, -wpy
	case 270:
		fallthrough
	case -90:
		return wpy, -wpx
	default:
		panic(deg)
	}
}

func executeInstruction2(x int, y int, wpx int, wpy int, ins instruction) (int, int, int, int) {

	switch ins.action {
	case "E":
		wpx += ins.amount
	case "W":
		wpx -= ins.amount
	case "N":
		wpy += ins.amount
	case "S":
		wpy -= ins.amount
	case "F":
		x += ins.amount * wpx
		y += ins.amount * wpy
	case "L":
		wpx, wpy = rotate2(ins.amount, wpx, wpy)
	case "R":
		wpx, wpy = rotate2(-ins.amount, wpx, wpy)
	}

	return x, y, wpx, wpy
}
