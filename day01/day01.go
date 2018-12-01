package main

import (
	"fmt"
	"log"
	"strconv"

	"../utils"
)

var seenFrequencies map[int]int
var currentFrequency int

func main() {
	currentFrequency = 0
	seenFrequencies = make(map[int]int)
	seenFrequencies[currentFrequency] = currentFrequency

	lines, err := utils.ReadLines("input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	// Part 1
	seenFrequencyTwice := parseFrequencies(lines)
	fmt.Printf("part 1: %d \n", currentFrequency)

	// Part 2
	for !seenFrequencyTwice {
		// loops until a frequency is seen twice
		seenFrequencyTwice = parseFrequencies(lines)
	}

	fmt.Printf("part 2: %d \n", currentFrequency)

}

func parseFrequencies(lines []string) bool {

	// parse every line of the input
	for _, line := range lines {

		deltaFrequency, _ := strconv.Atoi(line)
		currentFrequency += deltaFrequency

		// checks if the frequency has already been seen
		if _, ok := seenFrequencies[currentFrequency]; ok {
			return true
		}

		// adds the frequency to the map and increasing its value in the map
		seenFrequencies[currentFrequency]++

	}

	return false
}
