package main

import (
	"aoc_2023/common/utils"
	"testing"
)

func TestPart2(t *testing.T) {
	input, err := utils.ReadInputFromFile("./test_input.txt")
	if err != nil {
		t.Error(err.Error())
	}

	res, err := solution(input)
	if err != nil {
		t.Error(err)
	}

	expected := uint32(2286)
	if res != expected {
		t.Errorf("expected %d, got %d", expected, res)
	}
}
