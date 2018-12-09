package day09

import (
	"testing"

	"../utils"
	//	"../utils"
)

func TestMarbleScore(t *testing.T) {

	input1 := utils.ReadLines("testInput.txt")[0]
	testMarbleScore(t, input1, 32)

	input2 := utils.ReadLines("input.txt")[0]
	testMarbleScore(t, input2, 373597)

	input3 := "10 players; last marble is worth 1618 points"
	testMarbleScore(t, input3, 8317)

	input4 := "13 players; last marble is worth 7999 points"
	testMarbleScore(t, input4, 146373)

	input5 := "17 players; last marble is worth 1104 points"
	testMarbleScore(t, input5, 2764)

	input6 := "21 players; last marble is worth 6111 points"
	testMarbleScore(t, input6, 54718)

	input7 := "30 players; last marble is worth 5807 points"
	testMarbleScore(t, input7, 37305)

}

func testMarbleScore(t *testing.T, input string, expected int) {

	output := marbleScore(input)

	if output != expected {
		t.Errorf("marbleScore part 1 was incorrect, got: %d, want: %d.", output, expected)
	}
}
