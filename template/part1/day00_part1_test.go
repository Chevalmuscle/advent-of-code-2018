package templatepart1

import (
	"testing"

	"../../utils"
)

func TestSum(t *testing.T) {

	input1 := []string{"5", "30"}
	testSum(t, input1, 35)

	input2 := utils.ReadLines("../input.txt")
	testSum(t, input2, 8)

}

func testSum(t *testing.T, input []string, expected int) {

	output := sum(input)

	if output != expected {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", output, expected)
	}
}
