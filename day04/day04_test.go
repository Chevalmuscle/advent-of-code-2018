package day04part1

import (
	"testing"

	"../utils"
)

func TestGuardForSneaking(t *testing.T) {
	input1 := utils.ReadLines("testInput.txt")
	testGuardForSneaking(t, input1, 240, 4455)

	input2 := utils.ReadLines("input.txt")
	testGuardForSneaking(t, input2, 63509, 47910)
}

func testGuardForSneaking(t *testing.T, input []string, expected1 int, expected2 int) {
	output1, output2 := guardForSneaking(input)

	if output1 != expected1 {
		t.Errorf("Strategy 1 was incorrect, got: %d, want: %d.", output1, expected2)
	}
	if output2 != expected2 {
		t.Errorf("Strategy 2 was incorrect, got: %d, want: %d.", output2, expected2)
	}
}
