package day06

import (
	"testing"

	"../utils"
)

func TestSizeLargestArea(t *testing.T) {
	input1 := utils.ReadLines("testInput.txt")
	testSizeLargestArea(t, input1, 17)

	input2 := utils.ReadLines("input.txt")
	testSizeLargestArea(t, input2, 3569)
}

func testSizeLargestArea(t *testing.T, input []string, expected int) {
	output := sizeLargestArea(input)

	if output != expected {
		t.Errorf("sizeLargestArea was incorrect, got: %d, want: %d.", output, expected)
	}
}
