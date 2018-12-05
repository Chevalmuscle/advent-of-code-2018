package day05

import (
	"time"

	"../utils"
)

func getNeutralPolymer(polymer string) int {
	defer utils.TimeTaken(time.Now())

	var runes = []rune(polymer)

	for i := 1; i < len(runes); i++ {
		if isReacting(runes[i], runes[i-1]) {
			runes = append(runes[:i], runes[i+1:]...)
			runes = append(runes[:i-1], runes[i:]...)
			i -= 2
			if i < 0 {
				i = 0
			}
		}
	}

	return len(runes)
}

// difference of 32 between lower and upper case
func isReacting(unit1 rune, unit2 rune) bool {
	return (unit1-unit2) == 32 || (unit1-unit2) == -32
}
