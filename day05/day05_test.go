package day05

import (
	"testing"

	"../utils"
)

func TestNeutralPolymer(t *testing.T) {
	inputTest := "dabAcCaCBAcCcaDA"
	testNeutralPolymer(t, inputTest, 10)

	input := utils.ReadLines("input.txt")[0]
	testNeutralPolymer(t, input, 9900)
}

func testNeutralPolymer(t *testing.T, input string, expected int) {
	output := getNeutralPolymer(input)

	if output != expected {
		t.Errorf("NeutralPolymer was incorrect, got: %d, want: %d.", output, expected)
	}
}
