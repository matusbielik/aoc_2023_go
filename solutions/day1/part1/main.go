package main

import (
	"aoc_2023/common/utils"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	input, err := utils.ReadInputFromFile("./input.txt")
	if err != nil {
		log.Fatal(err.Error())
	}

	result := solution(input)
	fmt.Print(result)
}

func solution(input string) int {
	out := 0

	re := regexp.MustCompile(`\d`)

	lines := strings.Split(input, "\n")
	for _, line := range lines {

		digits := re.FindAllString(line, -1)
		digits = handleDigits(digits)
		if len(digits) < 1 {
			continue
		}

		digitsStr := strings.Join(digits, "")
		lineVal, err := strconv.Atoi(digitsStr)
		if err != nil {
			log.Fatal(err.Error())
		}

		out += lineVal
	}

	return out
}

func handleDigits(digits []string) []string {
	digitsLen := len(digits)
	if digitsLen < 1 {
		return nil
	} else if digitsLen == 1 {
		return []string{digits[0], digits[0]}
	} else if digitsLen > 2 {
		return []string{digits[0], digits[digitsLen-1]}
	}

	return digits
}
