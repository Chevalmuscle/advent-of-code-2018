package day08

import (
	"testing"

	"../utils"
)

func TestSumMetadata(t *testing.T) {

	input1 := utils.ReadLines("testInput.txt")[0]
	testSumMetadata(t, input1, 138, 66)

	input2 := utils.ReadLines("input.txt")[0]
	testSumMetadata(t, input2, 35911, 17206)

}

func testSumMetadata(t *testing.T, input string, expected1 int, expected2 int) {

	output1, output2 := sumMetadata(input)

	if output1 != expected1 {
		t.Errorf("sumMetadata part 1 was incorrect, got: %d, want: %d.", output1, expected1)
	}
	if output2 != expected2 {
		t.Errorf("sumMetadata part 2 was incorrect, got: %d, want: %d.", output2, expected2)
	}
}
