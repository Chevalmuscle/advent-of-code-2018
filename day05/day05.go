package day05

import (
	"strings"
	"time"

	"../utils"
)

func getNeutralPolymer(polymer string) (int, int) {
	defer utils.TimeTaken(time.Now())

	var units = []rune(polymer)
	var lowestLength = 99999999999
	var part1 = getLengthOfNeutralPolymer(units)
	var possibleUnits = getUniqueRunes(units)

	for lowerCase, upperCase := range possibleUnits {
		var tmp = strings.Replace(polymer, string(lowerCase), "", -1)
		tmp = strings.Replace(tmp, string(upperCase), "", -1)

		var currentLength = getLengthOfNeutralPolymer([]rune(tmp))

		if currentLength < lowestLength {
			lowestLength = currentLength
		}
	}

	return part1, lowestLength
}

// returns a map with the runes in the array
// key: rune in lowercase; value: rune in uppercase
// has to be only letters
func getUniqueRunes(runes []rune) map[rune]rune {
	var seenRunes = make(map[rune]rune)
	for _, currentRune := range runes {
		if currentRune > 96 { // lower cases
			if _, knownRune := seenRunes[currentRune]; !knownRune {
				seenRunes[currentRune] = currentRune - 32
			}
		}
	}
	return seenRunes

}

// difference of 32 between lower and upper case
func isReacting(unit1 rune, unit2 rune) bool {
	return (unit1-unit2) == 32 || (unit1-unit2) == -32
}

func getLengthOfNeutralPolymer(units []rune) int {
	for i := 1; i < len(units); i++ {
		if isReacting(units[i], units[i-1]) {
			units = append(units[:i], units[i+1:]...)
			units = append(units[:i-1], units[i:]...)
			i -= 2
			if i < 0 {
				i = 0
			}
		}
	}
	return len(units)
}
