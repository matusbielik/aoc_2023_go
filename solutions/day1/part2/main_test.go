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

	result := solution(input)

	expected := 281
	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}
