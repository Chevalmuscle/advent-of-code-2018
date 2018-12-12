package day12

import (
	"testing"

	"../utils"
)

func TestSumNumberWithPlants(t *testing.T) {
	expected1 := 325
	output1, _ := sumNumberWithPlants(utils.ReadLines("testInput.txt"))
	if output1 != expected1 {
		t.Errorf("sumNumberWithPlants part 1 with example data was incorrect, got: %d, want: %d.", output1, expected1)
	}

	input2 := utils.ReadLines("input.txt")
	testSumNumberWithPlants(t, input2, 6201, 9300000001023)
}

func testSumNumberWithPlants(t *testing.T, input []string, expected1 int, expected2 int) {
	output1, output2 := sumNumberWithPlants(input)

	if output1 != expected1 {
		t.Errorf("sumNumberWithPlants part 1 was incorrect, got: %d, want: %d.", output1, expected1)
	}
	if output2 != expected2 {
		t.Errorf("sumNumberWithPlants part 2 was incorrect, got: %d, want: %d.", output2, expected2)
	}
}
