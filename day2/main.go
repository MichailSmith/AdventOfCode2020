package main

import (
	"fmt"
	"regexp"
	"strconv"

	"adventofcode.com/misha/shared"
)

func main() {
	passwordStrings := shared.ReadFileLines("day2/input.txt")
	passwords := make([]passwordPolicy, len(passwordStrings))
	regex := regexp.MustCompile(`(\d+)-(\d+) (.): (.*)`)

	for i := range passwordStrings {
		values := regex.FindStringSubmatch(passwordStrings[i])

		min, err := strconv.Atoi(values[1])
		if err != nil {
			panic(err)
		}

		max, err := strconv.Atoi(values[2])
		if err != nil {
			panic(err)
		}

		passwords[i] = passwordPolicy{
			minCount:  min,
			maxCount:  max,
			character: values[3],
			password:  values[4],
		}
	}

	validCount := 0

	for i := range passwords {
		password := passwords[i]

		isValidRegex := regexp.MustCompile(password.character)

		count := len(isValidRegex.FindAllString(password.password, -1))

		if count >= password.minCount && count <= password.maxCount {
			validCount++
		}
	}

	fmt.Println(validCount)

	validCount = 0

	for i := range passwords {
		password := passwords[i]

		firstIndexContainsCharacter := password.password[password.minCount-1] == password.character[0]
		secondIndexContainsCharacter := password.password[password.maxCount-1] == password.character[0]

		if firstIndexContainsCharacter != secondIndexContainsCharacter {
			validCount++
		}
	}

	fmt.Println(validCount)
}

type passwordPolicy struct {
	minCount  int
	maxCount  int
	character string
	password  string
}
