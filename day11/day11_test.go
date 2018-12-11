package day11

import (
	"testing"
)

func TestCoordinates(t *testing.T) {

	input1 := 18
	testCoordinates(t, input1, "33,45")

	input2 := 42
	testCoordinates(t, input2, "21,61")

	finalInput := 1788
	testCoordinates(t, finalInput, "235,35")

}

func TestPowerLevel(t *testing.T) {

	testPowerLevel(t, 122, 79, 57, -5)
	testPowerLevel(t, 217, 196, 39, 0)
	testPowerLevel(t, 101, 153, 71, 4)

}

func testCoordinates(t *testing.T, input int, expected string) {

	output := coordinates(input)

	if output != expected {
		t.Errorf("Coordinates part 1 was incorrect, got: %v, want: %v.", output, expected)
	}
}

func testPowerLevel(t *testing.T, x int, y int, serialNumber int, expected int) {
	output := getPowerLevel(x, y, serialNumber)

	if output != expected {
		t.Errorf("getPowerLevel was incorrect, got: %v, want: %v.", output, expected)
	}
}
