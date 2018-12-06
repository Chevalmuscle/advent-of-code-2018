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

func TestIsReacting(t *testing.T) {
	testIsReacting(t, 65, 97, true)
	testIsReacting(t, 122, 90, true)
	testIsReacting(t, 103, 103, false)
	testIsReacting(t, 72, 72, false)
	testIsReacting(t, 118, 111, false)
	testIsReacting(t, 99, 80, false)
	testIsReacting(t, 81, 119, false)
	testIsReacting(t, 82, 84, false)
}

func testIsReacting(t *testing.T, input1 rune, input2 rune, expected bool) {

	var output = isReacting(input1, input2)

	if output != expected {
		t.Errorf("isReacting was incorrect, got: %t, want: %t.", output, expected)
	}
}

func testNeutralPolymer(t *testing.T, input string, expected1 int, expected2 int) {
	var output1, output2 = getNeutralPolymer(input)

	if output1 != expected1 {
		t.Errorf("Part 1 NeutralPolymer was incorrect, got: %d, want: %d.", output1, expected1)
	}
	if output2 != expected2 {
		t.Errorf("Part 2 Shortest neutral polymer was incorrect, got: %d, want: %d.", output2, expected2)
	}
}
