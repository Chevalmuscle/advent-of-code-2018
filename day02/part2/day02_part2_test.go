package main

import "testing"

func TestMain(t *testing.T) {
	main()
}

func TestHashMapWay(t *testing.T) {

	input1 := []string{"abcde", "fghij", "klmno", "pqrst", "fguij", "axcye", "wvxyz"}
	input2 := []string{"abcde", "aguij", "klmno", "pqrst", "fguij", "axcye", "wvxyz"}
	input3 := []string{"abcde", "fguij", "klmno", "pqrst", "fguij", "axcye", "wvxyz"}

	testHashMapWay(t, input1, "fghij", "fguij")
	testHashMapWay(t, input2, "aguij", "fguij")
	testHashMapWay(t, input3, "", "")

}

func TestNaiveWay(t *testing.T) {

	input1 := []string{"abcde", "fghij", "klmno", "pqrst", "fguij", "axcye", "wvxyz"}
	input2 := []string{"abcde", "aguij", "klmno", "pqrst", "fguij", "axcye", "wvxyz"}
	input3 := []string{"abcde", "fguij", "klmno", "pqrst", "fguij", "axcye", "wvxyz"}

	testNaiveWay(t, input1, "fghij", "fguij")
	testNaiveWay(t, input2, "aguij", "fguij")
	testNaiveWay(t, input3, "", "")

}

func TestExtractCommonLetters(t *testing.T) {

	commonLetters := extractCommonLetters("fghij", "fguij")

	if commonLetters != "fgij" {
		t.Errorf("commonLetters was incorrect, got: %s, want: %s.", commonLetters, "fgij")
	}
}

func TestsAlmostEqual(t *testing.T) {

	testAlmostEqual(t, "fghij", "fguij", 1, true)
	testAlmostEqual(t, "aaaa", "aaaa", 2, false)
	testAlmostEqual(t, "abab", "abab", 0, true)
	testAlmostEqual(t, "abab", "aaaa", 2, true)

}

func testHashMapWay(t *testing.T, input []string, expected1 string, expected2 string) {
	output1, output2 := hashMapWay(input)

	if output1 != expected1 {
		t.Errorf("output1 was incorrect, got: %s, want: %s.", output1, expected1)
	}
	if output2 != expected2 {
		t.Errorf("output2 was incorrect, got: %s, want: %s.", output2, expected2)
	}
}

func testNaiveWay(t *testing.T, input []string, expected1 string, expected2 string) {
	output1, output2 := naiveWay(input)

	if output1 != expected1 {
		t.Errorf("output1 was incorrect, got: %s, want: %s.", output1, expected1)
	}
	if output2 != expected2 {
		t.Errorf("output2 was incorrect, got: %s, want: %s.", output2, expected2)
	}
}

func testAlmostEqual(t *testing.T, str1 string, str2 string, allowedMistakeds int, expected bool) {

	output := almostEqual(str1, str2, allowedMistakeds)

	if output != expected {
		t.Errorf("got: %t, want: %t.", output, expected)
	}

}
