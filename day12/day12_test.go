package day12

import (
	"testing"

	"../utils"
)

func TestSumNumberWithPlants(t *testing.T) {

	input1 := utils.ReadLines("testInput.txt")
	testSumNumberWithPlants(t, input1, 325)

	input2 := utils.ReadLines("input.txt")
	testSumNumberWithPlants(t, input2, 6201)

}

func testSumNumberWithPlants(t *testing.T, input []string, expected int) {

	output := sumNumberWithPlants(input)

	if output != expected {
		t.Errorf("sumNumberWithPlants part 1 was incorrect, got: %d, want: %d.", output, expected)
	}
}
