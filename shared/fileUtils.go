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
