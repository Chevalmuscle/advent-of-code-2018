package day04part1

import (
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	"../utils"
)

type strategy2 struct {
	guard                int
	minute               int
	timeSleepOnTheMinute int
}

func guardForSneaking(records []string) (int, int) {
	defer utils.TimeTaken(time.Now())

	sort.Strings(records)

	var guardsSleepMinutes = make(map[int]map[int]int) // map[guardID]map[hour][timesSleptAtThatHour]
	var guardsSleepTime = make(map[int]int)
	var strategy2Guard = strategy2{guard: -1, minute: -1, timeSleepOnTheMinute: -1}

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
				guardsSleepMinutes[currentGuardID] = make(map[int]int)
			}

			if strings.Contains(record, "falls asleep") {
				sleepStartMinute = currentMinute

			} else if strings.Contains(record, "wakes up") {

				for minute := sleepStartMinute; minute < currentMinute; minute++ {
					guardsSleepTime[currentGuardID]++
					guardsSleepMinutes[currentGuardID][minute]++

					if guardsSleepMinutes[currentGuardID][minute] > strategy2Guard.timeSleepOnTheMinute {
						strategy2Guard = strategy2{
							guard:                currentGuardID,
							minute:               minute,
							timeSleepOnTheMinute: guardsSleepMinutes[currentGuardID][minute]}
					}
				}
				if guardsSleepTime[currentGuardID] > guardsSleepTime[sleepiestGuard] {
					sleepiestGuard = currentGuardID
				}
			}
		}
	}

	strategy1Minute := getKeyWithBiggerValue(guardsSleepMinutes[sleepiestGuard])
	strategy2Minute := strategy2Guard.minute

	return (sleepiestGuard * strategy1Minute), (strategy2Guard.guard * strategy2Minute)
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
