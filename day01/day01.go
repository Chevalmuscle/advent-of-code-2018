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

	seenFrequencyTwice := parseInput()
	fmt.Printf("part 1: %d \n", currentFrequency)

	for !seenFrequencyTwice {
		// loops until a frequency is seen twice
		seenFrequencyTwice = parseInput()
	}

}

func parseInput() bool {
	lines, err := utils.ReadLines("input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	// parse every line of the input
	for _, line := range lines {

		deltaFrequency, _ := strconv.Atoi(line)
		currentFrequency += deltaFrequency

		// checks if the frequency has already been seen
		if _, ok := seenFrequencies[currentFrequency]; ok {
			fmt.Printf("part 2: %d\n", currentFrequency)
			return true
		}

		// adds the frequency to the map and increasing its value in the map
		seenFrequencies[currentFrequency]++

	}

	return false
}
