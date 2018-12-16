package day14

import (
	"testing"
)

func TestReceipes(t *testing.T) {

	testReceipes(t, 9, "5158916779")
	testReceipes(t, 5, "0124515891")
	testReceipes(t, 18, "9251071085")
	testReceipes(t, 2018, "5941429882")

	finalInput := 768071
	testReceipes(t, finalInput, "6548103910")
}

func testReceipes(t *testing.T, input int, expected string) {

	output := recipes(input)

	if output != expected {
		t.Errorf("testReceipes part 1 was incorrect, got: %v, want: %v", output, expected)
	}
}
