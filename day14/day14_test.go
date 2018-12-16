package day14

import (
	"testing"
)

func TestReceipes(t *testing.T) {
	testReceipesPart1(t, 9, "5158916779")
	testReceipesPart1(t, 18, "9251071085")
	testReceipesPart1(t, 2018, "5941429882")

	testReceipesPart2(t, 51589, 9)
	testReceipesPart2(t, 92510, 18)
	testReceipesPart2(t, 59414, 2018)

	// my input
	testReceipesPart1(t, 768071, "6548103910")
	testReceipesPart2(t, 768071, 20198090)
}

func testReceipesPart1(t *testing.T, input int, expected string) {
	output := recipesPart1(input)

	if output != expected {
		t.Errorf("testReceipes part 1 was incorrect, got: %v, want: %v", output, expected)
	}
}

func testReceipesPart2(t *testing.T, input int, expected int) {
	output := recipesPart2(input)

	if output != expected {
		t.Errorf("testReceipes part 2 was incorrect, got: %v, want: %v", output, expected)
	}
}
