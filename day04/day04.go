package day04part1

import (
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	"../utils"
)

type guard struct {
	id                   int
	minute               int
	timeSleepOnTheMinute int
}

type sleepTimes struct {
	total    int
	byMinute map[int]int
}

func guardForSneaking(records []string) (int, int) {
	defer utils.TimeTaken(time.Now())

	sort.Strings(records)

	var guardsSleepMinutes = make(map[int]sleepTimes)
	var mostConsistentGuard = guard{id: -1, minute: -1, timeSleepOnTheMinute: -1}

	var currentGuardID int
	var sleepStartMinute int
	var sleepiestGuard int

	for _, record := range records {
		var currentMinute, _ = strconv.Atoi(getMinute(record))
		if strings.ContainsAny(record, "#") {
			// a new guard begins it's shift
			currentGuardID = getGuardID(record)
		} else {
			if _, knownGuard := guardsSleepMinutes[currentGuardID]; !knownGuard {
				guardsSleepMinutes[currentGuardID] = sleepTimes{total: 0, byMinute: make(map[int]int)}
			}
			var currentGuard = guardsSleepMinutes[currentGuardID]

			if strings.Contains(record, "falls asleep") {
				sleepStartMinute = currentMinute

			} else if strings.Contains(record, "wakes up") {

				for minute := sleepStartMinute; minute < currentMinute; minute++ {
					// because it can't be pass by ref..
					guardsSleepMinutes[currentGuardID] = sleepTimes{total: guardsSleepMinutes[currentGuardID].total + 1, byMinute: guardsSleepMinutes[currentGuardID].byMinute}
					currentGuard.byMinute[minute]++

					if currentGuard.byMinute[minute] > mostConsistentGuard.timeSleepOnTheMinute {
						mostConsistentGuard = guard{
							id:                   currentGuardID,
							minute:               minute,
							timeSleepOnTheMinute: currentGuard.byMinute[minute]}
					}
				}
				if currentGuard.total > guardsSleepMinutes[sleepiestGuard].total {
					sleepiestGuard = currentGuardID
				}
			}
		}
	}

	strategy1Minute := getKeyWithBiggerValue(guardsSleepMinutes[sleepiestGuard].byMinute)
	strategy2Minute := mostConsistentGuard.minute

	return (sleepiestGuard * strategy1Minute), (mostConsistentGuard.id * strategy2Minute)
}

func getMinute(record string) string {
	return regexp.MustCompile("([0-9]{2})").FindAllString(record, -1)[5]
}
func getGuardID(record string) int {
	var guardID, _ = strconv.Atoi(strings.Replace(regexp.MustCompile("#([0-9]+)").FindAllString(record, -1)[0], "#", "", -1))
	return guardID
}

// has to change (cannot handle negative values)
func getKeyWithBiggerValue(m map[int]int) int {
	var k = -1
	var biggestValue = -1

	for key, value := range m {
		if value > biggestValue {
			biggestValue = value
			k = key
		}
	}
	return k
}
