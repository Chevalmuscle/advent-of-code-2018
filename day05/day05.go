package day05

import (
	"strings"
	"sync"
	"time"

	"../utils"
)

var wg = sync.WaitGroup{}
var minPolymerLengthforGoroutine int

func getNeutralPolymerLength(polymer string) (int, int) {
	defer utils.TimeTaken(time.Now())

	minPolymerLengthforGoroutine = len(polymer)/8 - 1
	var units = []rune(polymer)

	// part1
	var channel = make(chan []rune, 1)
	wg.Add(1)
	go routineRecursiveGetNeutralPolymer(channel, units)
	wg.Wait()
	var neutralPolymer = <-channel

	var lengthInitialNeutralPolymer = len(neutralPolymer)

	//part 2
	var uniqueUnits = getUniqueRunes(units)
	var lowestLength = lengthInitialNeutralPolymer

	for lowerCaseUnit, upperCaseUnit := range uniqueUnits {

		var optimizedPolymer = strings.Replace(polymer, string(lowerCaseUnit), "", -1)
		optimizedPolymer = strings.Replace(optimizedPolymer, string(upperCaseUnit), "", -1)

		wg.Add(1)
		go routineRecursiveGetNeutralPolymer(channel, []rune(optimizedPolymer))
		wg.Wait()
		var neutralPolymer = <-channel

		var currentLength = len(neutralPolymer)

		if currentLength < lowestLength {
			lowestLength = currentLength
		}
	}
	return lengthInitialNeutralPolymer, lowestLength
}

func routineRecursiveGetNeutralPolymer(ch chan<- []rune, units []rune) {
	if len(units) > minPolymerLengthforGoroutine && len(units) > 1 {
		var channel1 = make(chan []rune, 1)
		var channel2 = make(chan []rune, 1)
		wg.Add(2)

		var firstHalf = units[:len(units)/2]
		go routineRecursiveGetNeutralPolymer(channel1, firstHalf)

		var secondHalf = units[len(units)/2:]
		go routineRecursiveGetNeutralPolymer(channel2, secondHalf)

		//wg.Wait()
		var firstHalfNeutral = <-channel1
		var secondHalfNeutral = <-channel2
		var mergedPolymer = append(firstHalfNeutral, secondHalfNeutral...)
		var neutralPolymer = getNeutralPolymer(mergedPolymer)

		ch <- neutralPolymer

	} else {
		ch <- recursiveGetNeutralPolymer(units)
	}
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
