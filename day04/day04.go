package day04part1

import (
	"math"
	"reflect"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	"../utils"
)

type date struct {
	Year  int
	Month int
	Day   int
	Hour  int
	Min   int
}

type strategy2 struct {
	guard                string
	minute               int
	timeSleepOnTheMinute int
}

var wakesEvent = "wakes"
var fallsAsleepEvent = "asleep"
var beginShiftEvent = "#"

func guardForSneaking(unsortedRecords []string) (int, int) {
	defer utils.TimeTaken(time.Now())

	var sortedRecords = make([]date, len(unsortedRecords))
	var eventsMap = make(map[date]string)
	var guardsWorkHours = make(map[string]map[int]int) // map[guardID]map[hour][timesWorkedAtThatHour]
	var guardsSleepTime = make(map[string]int)

	var strategy2Guard = strategy2{guard: "#-1", minute: -1, timeSleepOnTheMinute: -1}

	for i, line := range unsortedRecords {
		var date = getDate(line)
		sortedRecords[i] = date
		var event = getEvent(line)
		eventsMap[date] = event
	}

	sort.SliceStable(sortedRecords, func(i, j int) bool {
		return sortedRecords[i].compareTo(sortedRecords[j]) == -1
	})

	var currentGuardID string //= eventsMap[sortedRecords[0]]
	var sleepStartTime = sortedRecords[0]
	var sleepiestGuard = eventsMap[sleepStartTime]

	for _, currentDate := range sortedRecords {

		if strings.ContainsAny(eventsMap[currentDate], beginShiftEvent) {
			// a new guard begins it's shift
			currentGuardID = eventsMap[currentDate]
		} else {
			if _, knownGuard := guardsWorkHours[currentGuardID]; !knownGuard {
				guardsWorkHours[currentGuardID] = make(map[int]int)
			}

			if eventsMap[currentDate] == fallsAsleepEvent {
				sleepStartTime = currentDate

			} else if eventsMap[currentDate] == wakesEvent {
				guardsSleepTime[currentGuardID] += currentDate.minutesSince(sleepStartTime)

				for minute := sleepStartTime.Min; minute < currentDate.Min; minute++ {
					guardsWorkHours[currentGuardID][minute]++

					if guardsWorkHours[currentGuardID][minute] > strategy2Guard.timeSleepOnTheMinute {
						strategy2Guard = strategy2{guard: currentGuardID, minute: minute, timeSleepOnTheMinute: guardsWorkHours[currentGuardID][minute]}
					}
				}

				if guardsSleepTime[currentGuardID] > guardsSleepTime[sleepiestGuard] {
					sleepiestGuard = currentGuardID
				}
			}
		}
	}

	strategy1GuardID, _ := strconv.Atoi(strings.Replace(sleepiestGuard, "#", "", -1))
	strategy1Minute := getKeyWithBiggerValue(guardsWorkHours[sleepiestGuard])

	strategy2GuardID, _ := strconv.Atoi(strings.Replace(strategy2Guard.guard, "#", "", -1))
	strategy2Minute := strategy2Guard.minute

	return (strategy1GuardID * strategy1Minute), (strategy2GuardID * strategy2Minute)
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

func getEvent(record string) string {
	return regexp.MustCompile("(asleep|wakes|#[0-9]+)").FindAllString(record, -1)[0]
}

func getDate(record string) date {
	var re = regexp.MustCompile("([0-9]+)")

	var year, _ = strconv.Atoi(re.FindAllString(record, -1)[0])
	var month, _ = strconv.Atoi(re.FindAllString(record, -1)[1])
	var day, _ = strconv.Atoi(re.FindAllString(record, -1)[2])
	var hour, _ = strconv.Atoi(re.FindAllString(record, -1)[3])
	var min, _ = strconv.Atoi(re.FindAllString(record, -1)[4])

	return date{Year: year, Month: month, Day: day, Hour: hour, Min: min}
}

func (date1 date) minutesSince(date2 date) int {
	return date1.Min - date2.Min
}

// if date1 > date2, returns 1;
// if date1 < date2, returns -1;
// if date1 = date2, returns 0
func (date1 date) compareTo(date2 date) int {
	var struct1 = reflect.ValueOf(&date1).Elem()
	var struct2 = reflect.ValueOf(&date2).Elem()

	for i := 0; i < struct1.NumField(); i++ {
		var difference = struct1.Field(i).Interface().(int) - struct2.Field(i).Interface().(int)
		if difference != 0 {
			return difference / int(math.Abs(float64(difference)))
		}
	}
	return 0
}
