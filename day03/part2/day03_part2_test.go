package day03part2

import (
	"testing"

	"../../utils"
)

func TestCleanClaimID(t *testing.T) {

	input1 := []string{"#1 @ 1,3: 4x4", "#2 @ 3,1: 4x4", "#3 @ 5,5: 2x2"}
	testCleanClaimID(t, input1, 3)

	input2 := utils.ReadLines("../input.txt")
	testCleanClaimID(t, input2, 504)

}

func testCleanClaimID(t *testing.T, input []string, expected int) {

	output := getCleanClaimID(input)

	if output != expected {
		t.Errorf("countOverLap was incorrect, got: %d, want: %d.", output, expected)
	}
}
