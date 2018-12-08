package day06part2

import (
	"testing"

	"../../utils"
)

func TestSizeLargestArea(t *testing.T) {
	input11 := utils.ReadLines("../testInput.txt")
	input12 := 32
	testSizeLargestArea(t, input11, input12, 16)

	input21 := utils.ReadLines("../input.txt")
	input22 := 10000
	testSizeLargestArea(t, input21, input22, 48978)
}

func testSizeLargestArea(t *testing.T, input1 []string, input2 int, expected int) {
	output := sizeLargestArea(input1, input2)

	if output != expected {
		t.Errorf("sizeLargestArea was incorrect, got: %d, want: %d.", output, expected)
	}
}
