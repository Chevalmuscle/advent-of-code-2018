package day03

import (
	"testing"

	"../../utils"
	//	"../utils"
)

func TestCountOverlap(t *testing.T) {

	input1 := []string{"#1 @ 1,3: 4x4", "#2 @ 3,1: 4x4", "#3 @ 5,5: 2x2"}
	testCountOverlap(t, input1, 4)

	input2 := utils.ReadLines("../input.txt")
	testCountOverlap(t, input2, 109785)

}

func testCountOverlap(t *testing.T, input []string, expected int) {

	output := countOverlap(input)

	if output != expected {
		t.Errorf("countOverLap was incorrect, got: %d, want: %d.", output, expected)
	}
}
