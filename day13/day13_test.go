package day13

import (
	"testing"

	"../utils"
)

func TestFirstCrash(t *testing.T) {
	/*
		input1 := utils.ReadLines("testInput1.txt")
		expected1 := "7,3"
		output1, _ := firstCrash(input1)
		if output1 != expected1 {
			t.Errorf("Crash part 1 was incorrect, got: %v, want: %v", output1, expected1)
		}

		input2 := utils.ReadLines("testInput2.txt")
		expected2 := "6,4"
		_, output2 := firstCrash(input2)
		if output2 != expected2 {
			t.Errorf("Crash part 2 was incorrect, got: %v, want: %v", output2, expected2)
		}
	*/
	input3 := utils.ReadLines("input.txt")
	testFirstCrash(t, input3, "124,90", "145,88") //not 42,109

}

func testFirstCrash(t *testing.T, input []string, expected1 string, expected2 string) {

	output1, output2 := firstCrash(input)

	if output1 != expected1 {
		t.Errorf("Crash part 1 was incorrect, got: %v, want: %v", output1, expected1)
	}
	if output2 != expected2 {
		t.Errorf("Crash part 2 was incorrect, got: %v, want: %v", output2, expected2)
	}
}
