package day08

import (
	"testing"

	"../utils"
)

func TestSumMetadata(t *testing.T) {

	input1 := utils.ReadLines("testInput.txt")[0]
	testSumMetadata(t, input1, 138)

	input2 := utils.ReadLines("input.txt")[0]
	testSumMetadata(t, input2, 35911)

}

func testSumMetadata(t *testing.T, input string, expected int) {

	output := sumMetadata(input)

	if output != expected {
		t.Errorf("sumMetadata part 1 was incorrect, got: %d, want: %d.", output, expected)
	}
}
