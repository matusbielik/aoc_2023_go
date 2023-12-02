package main

import (
	"aoc_2023/common/utils"
	"fmt"
	"log"
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

func solution(in string) (int, error) {
	out := 0

	return out, nil
}
