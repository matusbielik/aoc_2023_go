package main

import (
	"aoc_2023/common/utils"
	"aoc_2023/solutions/day2"
	"errors"
	"fmt"
	"log"
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

var errInvalidInput = errors.New("invalid input")

func solution(in string) (uint32, error) {
	var out uint32

	games := strings.Split(in, "\n")
	for _, game := range games {
		_, allHands, found := strings.Cut(game, ":")
		if !found {
			return 0, errInvalidInput
		}

		minimal := day2.CubesHand{}

		hands := strings.Split(allHands, ";")
		for _, handStr := range hands {
			hand := day2.NewHandFromString(handStr)
			minimal.LogMinimalVal(hand)
		}

		minimalPower := minimal.Power()
		out += minimalPower
	}

	return out, nil
}
