package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"adventofcode.com/misha/shared"
)

func main() {
	passportLines := shared.ReadFileLines("day4/input.txt")

	passports := []passport{}

	currentPassport := passport{
		fields: make(map[string]string),
	}
	for i := range passportLines {
		if len(passportLines[i]) == 0 {
			passports = append(passports, currentPassport)
			currentPassport = passport{
				fields: make(map[string]string),
			}
		} else {
			fields := strings.Split(passportLines[i], " ")

			for j := range fields {
				kvp := strings.Split(fields[j], ":")
				if len(kvp) == 2 {
					currentPassport.fields[kvp[0]] = kvp[1]
				}
			}
		}
	}

	validators := map[string]func(string) bool{
		"byr": validByr,
		"iyr": validIyr,
		"eyr": validEyr,
		"hgt": validHgt,
		"hcl": validHcl,
		"ecl": validEcl,
		"pid": validPid,
	}

	validCount := 0
	completeCount := 0

	for i := range passports {

		if passports[i].isValid(validators) {
			validCount++
		}
		if passports[i].isComplete(validators) {
			completeCount++
		}
	}

	fmt.Println(validCount)
	fmt.Println(completeCount)
}

func isNumberInRange(value string, min int, max int) bool {
	val, err := strconv.Atoi(value)
	if err != nil {
		return false
	}

	return val >= min && val <= max
}

func validByr(byr string) bool {
	return isNumberInRange(byr, 1920, 2002)
}

func validIyr(iyr string) bool {
	return isNumberInRange(iyr, 2010, 2020)
}

func validEyr(eyr string) bool {
	return isNumberInRange(eyr, 2020, 2030)
}

func validHgt(hgt string) bool {
	index := (len(hgt) - 2)
	unit := hgt[index:]
	if unit == "in" {
		return isNumberInRange(hgt[:index], 59, 76)
	}
	if unit == "cm" {
		return isNumberInRange(hgt[:index], 150, 193)
	}
	return false
}

func validHcl(hcl string) bool {
	hclRegex := regexp.MustCompile(`^#[\da-f]{6}$`)
	return hclRegex.MatchString(hcl)
}

func validEcl(ecl string) bool {
	eclRegex := regexp.MustCompile(`^(amb)|(blu)|(brn)|(gry)|(grn)|(hzl)|(oth)$`)
	return eclRegex.MatchString(ecl)
}

func validPid(pid string) bool {
	pidRegex := regexp.MustCompile(`^\d{9}$`)
	return pidRegex.MatchString(pid)
}

type passport struct {
	fields map[string]string
}

func (p passport) isValid(validators map[string]func(string) bool) bool {
	if !p.isComplete(validators) {
		return false
	}
	for key, validator := range validators {
		if !validator(p.fields[key]) {
			return false
		}
	}
	return true
}

func (p passport) isComplete(validators map[string]func(string) bool) bool {
	for key := range validators {
		_, exists := p.fields[key]
		if !exists {
			return false
		}
	}
	return true
}
