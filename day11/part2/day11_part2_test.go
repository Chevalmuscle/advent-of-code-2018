package day11part2

import (
	"testing"
)

func TestCoordinates(t *testing.T) {
	testCoordinates(t, 18, "90,269,16")
	testCoordinates(t, 42, "232,251,12")
	testCoordinates(t, 1788, "142,265,7")
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
