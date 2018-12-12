package day12

import (
	"time"

	"../utils"
)

var sizeRule = 5

func sumNumberWithPlants(input []string) int {
	defer utils.TimeTaken(time.Now())

	var startBuffer = ".........."
	var state = startBuffer + input[0][15:] + "......................."
	var rules = make(map[string]string)

	for i := 2; i < len(input); i++ {
		rules[input[i][:5]] = input[i][9:]
	}

	for gen := 1; gen <= 20; gen++ {
		var newState string

		for i := 0; i < len(state); i++ {

			var scope string
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
		state = newState
	}

	var count = 0
	for i := 0; i < len(state); i++ {
		if string(state[i]) == "#" {
			count += i - len(startBuffer)
		}
	}

	return count
}
