package main

import (
	"aoc_2023/common/utils"
	"testing"
)

func TestPart1(t *testing.T) {
	input, err := utils.ReadInputFromFile("./test_input.txt")
	if err != nil {
		t.Error(err.Error())
	}

	expected := 0
	res, err := solution(input)
	if err != nil {
		t.Error(err)
	}

	if res != expected {
		t.Errorf("expected %d, got %d", expected, res)
	}
}

func BenchmarkPart1(b *testing.B) {
	inputStr, err := utils.ReadInputFromFile("./input.txt")
	if err != nil {
		b.Error(err.Error())
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		solution(inputStr)
	}
}
