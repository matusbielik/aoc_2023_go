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

	fmt.Print(solution(input))
}

var numbersMap map[string]string = map[string]string{
	"one":   "one1one",
	"two":   "two2two",
	"three": "three3three",
	"four":  "four4four",
	"five":  "five5five",
	"six":   "six6six",
	"seven": "seven7seven",
	"eight": "eight8eight",
	"nine":  "nine9nine",
}

func solution(input string) int {
	out := 0

	re := regexp.MustCompile(`\d`)

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		line = replaceWordsByDigits(line)
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

func replaceWordsByDigits(in string) string {
	for word, digit := range numbersMap {
		in = strings.ReplaceAll(in, word, digit)
	}

	return in
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
