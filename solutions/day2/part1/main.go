package main

import (
	aocErrors "aoc_2023/common/errors"

	"aoc_2023/common/utils"
	"aoc_2023/solutions/day2"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	input, err := utils.ReadInputFromFile("./input.txt")
	if err != nil {
		log.Fatal(err.Error())
	}

	res, err := solution(input)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("result is %d\n", res)
}

const (
	R day2.Color = "red"
	G day2.Color = "green"
	B day2.Color = "blue"
)

var rulesMax = day2.CubesHand{
	R: 12,
	G: 13,
	B: 14,
}

func solution(in string) (int, error) {
	out := 0

	lines := strings.Split(in, "\n")
	for _, line := range lines {
		header, draws_, found := strings.Cut(line, ":")
		if !found {
			return 0, aocErrors.ErrInvalidInput
		}

		gameId_ := strings.Replace(header, "Game ", "", -1)
		gameId, err := strconv.Atoi(gameId_)
		if err != nil {
			return 0, aocErrors.ErrInvalidInput
		}

		var allHandsValid = true
		hands := strings.Split(draws_, ";")
		for _, hand_ := range hands {
			hand := day2.NewHandFromString(hand_)
			if !hand.IsValid(rulesMax) {
				allHandsValid = false
				break
			}
		}

		if allHandsValid {
			out += gameId
		}
	}

	return out, nil
}
