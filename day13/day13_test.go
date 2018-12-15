package day13

import (
	"testing"

	"../utils"
)

func TestFirstCrash(t *testing.T) {

	input1 := utils.ReadLines("testInput.txt")
	testFirstCrash(t, input1, "7,3")

	input2 := utils.ReadLines("input.txt")
	testFirstCrash(t, input2, "lol")

}

func testFirstCrash(t *testing.T, input []string, expected string) {

	output := firstCrash(input)

	if output != expected {
		t.Errorf("Sum was incorrect, got: %v, want: %v", output, expected)
	}
}
