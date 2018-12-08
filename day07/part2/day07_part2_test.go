package day07

import (
	"testing"

	"../../utils"
)

// 183
func TestInstructionOrder(t *testing.T) {

	input21 := utils.ReadLines("../input.txt")
	input22 := 5
	testInstructionOrder(t, input21, input22, 903)

}

func testInstructionOrder(t *testing.T, input1 []string, input2 int, expected int) {
	output := instructionOrder(input1, input2)

	if output != expected {
		t.Errorf("instructionOrder part 2 was incorrect, got: %v, want: %v.", output, expected)
	}
}
