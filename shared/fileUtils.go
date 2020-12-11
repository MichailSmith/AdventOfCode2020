package shared

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func ReadFileLines(file string) []string {
	data, err := ioutil.ReadFile(file)

	if err != nil {
		fmt.Println("Error opening file", err)
	}

	return strings.Split(string(data), "\n")
}

func ReadFileBytes(file string) [][]byte {
	data, err := ioutil.ReadFile(file)

	if err != nil {
		fmt.Println("Error opening file", err)
	}
	result := [][]byte{}

	startIndex := 0
	for i, value := range data {
		if value == '\n' && i > startIndex {
			result = append(result, data[startIndex:i])
			startIndex = i + 1
		}
	}

	return result
}
