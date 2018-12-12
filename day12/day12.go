package day12

import (
	"time"

	"../utils"
)

var ruleSize = 5

func sumNumberWithPlants(input []string) (int, int) {
	defer utils.TimeTaken(time.Now())

	var borderBuffer = "..."
	var startBufferSize = len(borderBuffer)
	var state = borderBuffer + input[0][15:] + borderBuffer
	var potsGen20 int // part 1 answer

	var currentPotCount = 0
	var previousPotCount = 0
	var diff int

	var rules = make(map[string]string)
	for i := 2; i < len(input); i++ {
		rules[input[i][:ruleSize]] = input[i][ruleSize+4:]
	}

	var currentGen = 1
	for ; currentGen <= 200; currentGen++ {
		var newState string

		for i := 0; i < len(state); i++ {
			var scope string

			// to handle borders
			if i == 0 {
				scope = ".." + state[i:i+3]
			} else if i == 1 {
				scope = "." + state[i-1:i+3]
			} else if i == len(state)-1 {
				scope = state[i-2:len(state)-1] + ".."
			} else if i == len(state)-2 {
				scope = state[i-2:len(state)-2] + "."
			} else {
				scope = state[i-2 : i+3]
			}

			if _, isARule := rules[scope]; isARule {
				newState += rules[scope]
			} else {
				newState += "."
			}
		}

		currentPotCount = countPots(newState, startBufferSize)
		diff = currentPotCount - previousPotCount
		state = ".." + newState + ".."
		startBufferSize += 2
		previousPotCount = currentPotCount

		// for part 1
		if currentGen == 20 {
			potsGen20 = countPots(state, startBufferSize)
		}
	}
	// part 2
	potsGen50000000000 := currentPotCount + diff + diff*(50000000000-currentGen)
	return potsGen20, potsGen50000000000
}

func countPots(state string, lengthStartBuffer int) int {
	var count = 0
	for i := 0; i < len(state); i++ {
		if string(state[i]) == "#" {
			count += i - lengthStartBuffer
		}
	}
	return count
}
