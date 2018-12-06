package day05

import (
	"testing"

	"../utils"
)

func TestNeutralPolymer(t *testing.T) {
	inputTest := "dabAcCaCBAcCcaDA"
	testNeutralPolymer(t, inputTest, 10, 4)

	input := utils.ReadLines("input.txt")[0]
	testNeutralPolymer(t, input, 9900, 4992)

}

func testNeutralPolymer(t *testing.T, input string, expected1 int, expected2 int) {
	output1, output2 := getNeutralPolymer(input)

	if output1 != expected1 {
		t.Errorf("Part 1 NeutralPolymer was incorrect, got: %d, want: %d.", output1, expected1)
	}
	if output2 != expected2 {
		t.Errorf("Part 2 Shortest neutral polymer was incorrect, got: %d, want: %d.", output2, expected2)
	}
}
