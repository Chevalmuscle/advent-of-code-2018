package day05

import (
	"strings"
	"time"

	"../utils"
)

func getNeutralPolymerLength(polymer string) (int, int) {
	defer utils.TimeTaken(time.Now())

	var units = []rune(polymer)

	// part1
	var lengthInitialPolymer = len(recursiveGetNeutralPolymer(units))

	//part 2
	var uniqueUnits = getUniqueRunes(units)
	var lowestLength = lengthInitialPolymer

	for lowerCaseUnit, upperCaseUnit := range uniqueUnits {
		var polymerWithoutAUnit = strings.Replace(polymer, string(lowerCaseUnit), "", -1)
		polymerWithoutAUnit = strings.Replace(polymerWithoutAUnit, string(upperCaseUnit), "", -1)

		var currentLength = len(recursiveGetNeutralPolymer([]rune(polymerWithoutAUnit)))

		if currentLength < lowestLength {
			lowestLength = currentLength
		}
	}
	return lengthInitialPolymer, lowestLength
}

func recursiveGetNeutralPolymer(units []rune) []rune {
	if len(units) < 2 {
		return units
	}
	var firstHalf = units[:len(units)/2]
	var neutralPolymer1 = recursiveGetNeutralPolymer(firstHalf)

	var secondHalf = units[len(units)/2:]
	var neutralPolymer2 = recursiveGetNeutralPolymer(secondHalf)

	var mergedPolymer = append(neutralPolymer1, neutralPolymer2...)
	return getNeutralPolymer(mergedPolymer)
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

func getNeutralPolymer(units []rune) []rune {
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
	return units
}
