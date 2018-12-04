package day04part1

import (
	"testing"

	"../../utils"
	//"../../utils"
)

//"../../utils"

func TestGuardStrategy1(t *testing.T) {
	input1 := utils.ReadLines("../testInput.txt")
	testGuardStrategy1(t, input1, 240)

	input2 := utils.ReadLines("../input.txt")
	testGuardStrategy1(t, input2, 63509)
}
func TestGetDate(t *testing.T) {
	var input1 = "[1518-09-23 00:50] wakes up"
	var input2 = "[1518-05-14 00:27] falls asleep"
	var input3 = "[1518-04-18 00:00] Guard #2927 begins shift"

	testgetDate(t, input1, date{Year: 1518, Month: 9, Day: 23, Hour: 0, Min: 50})
	testgetDate(t, input2, date{Year: 1518, Month: 5, Day: 14, Hour: 0, Min: 27})
	testgetDate(t, input3, date{Year: 1518, Month: 4, Day: 18, Hour: 0, Min: 0})
}

func TestDateCompareTo(t *testing.T) {
	var date1 = date{Year: 1518, Month: 9, Day: 14, Hour: 0, Min: 50}
	var date2 = date{Year: 1518, Month: 5, Day: 14, Hour: 0, Min: 27}

	var greaterOutput = date1.compareTo(date2)
	var smallerOutput = date2.compareTo(date1)
	var equalsOutput = date1.compareTo(date1)

	if greaterOutput != 1 {
		t.Errorf("date.compareTo (greater) was incorrect, got: %d, want: %d.", greaterOutput, 1)
	}
	if smallerOutput != -1 {
		t.Errorf("date.compareTo (smaller) was incorrect, got: %d, want: %d.", smallerOutput, -1)
	}
	if equalsOutput != 0 {
		t.Errorf("date.compareTo (equal) was incorrect, got: %d, want: %d.", equalsOutput, 0)
	}
}

func testgetDate(t *testing.T, input string, expected date) {
	output := getDate(input)

	if output.compareTo(expected) != 0 {
		t.Errorf("Date was incorrect, got: %d, want: %d.", output, expected)
	}
}

func testGuardStrategy1(t *testing.T, input []string, expected int) {
	output := guardStrategy1(input)

	if output != expected {
		t.Errorf("Strategy 1 was incorrect, got: %d, want: %d.", output, expected)
	}
}
