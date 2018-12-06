package day05

import (
	"strings"
	"sync"
	"time"

	"../utils"
)

var wg = sync.WaitGroup{}

func getNeutralPolymerLength(polymer string) (int, int) {
	defer utils.TimeTaken(time.Now())

	var units = []rune(polymer)

	var channel1 = make(chan []rune, 1)
	var channel2 = make(chan []rune, 1)
	wg.Add(2)
	var firstHalf = units[:len(units)/2]
	var secondHalf = units[len(units)/2:]
	go routineGetNeutralPolymer(channel1, firstHalf)
	go routineGetNeutralPolymer(channel2, secondHalf)

	wg.Wait()
	var firstHalfNeutral = <-channel1
	var secondHalfNeutral = <-channel2
	var mergedPolymer = append(firstHalfNeutral, secondHalfNeutral...)
	var neutralPolymer = getNeutralPolymer(mergedPolymer)

	// part1
	var lengthInitialNeutralPolymer = len(neutralPolymer)

	//part 2
	var uniqueUnits = getUniqueRunes(units)
	var lowestLength = lengthInitialNeutralPolymer

	for lowerCaseUnit, upperCaseUnit := range uniqueUnits {

		var polymerWithoutAUnit = strings.Replace(polymer, string(lowerCaseUnit), "", -1)
		polymerWithoutAUnit = strings.Replace(polymerWithoutAUnit, string(upperCaseUnit), "", -1)

		wg.Add(2)

		var firstHalf = []rune(polymerWithoutAUnit)[:len(polymerWithoutAUnit)/2]
		var secondHalf = []rune(polymerWithoutAUnit)[len(polymerWithoutAUnit)/2:]
		go routineGetNeutralPolymer(channel1, firstHalf)
		go routineGetNeutralPolymer(channel2, secondHalf)

		wg.Wait()
		var firstHalfNeutral = <-channel1
		var secondHalfNeutral = <-channel2
		var mergedPolymer = append(firstHalfNeutral, secondHalfNeutral...)
		var neutralPolymer = getNeutralPolymer(mergedPolymer)

		var currentLength = len(neutralPolymer)

		if currentLength < lowestLength {
			lowestLength = currentLength
		}
	}
	return lengthInitialNeutralPolymer, lowestLength
}

func routineGetNeutralPolymer(ch chan<- []rune, units []rune) {
	ch <- recursiveGetNeutralPolymer(units)
	wg.Done()
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
