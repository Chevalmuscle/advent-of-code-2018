package day07part1

import (
	"testing"

	"../../utils"
)

// 183
func TestInstructionOrder(t *testing.T) {

	input1 := utils.ReadLines("../testInput.txt")
	testInstructionOrder(t, input1, "CABDFE")

	input2 := utils.ReadLines("../input.txt")
	testInstructionOrder(t, input2, "GKRVWBESYAMZDPTIUCFXQJLHNO")

}

func testInstructionOrder(t *testing.T, input []string, expected string) {
	output := instructionOrder(input)

	if output != expected {
		t.Errorf("instructionOrder part 1 was incorrect, got: %v, want: %v.", output, expected)
	}
}
