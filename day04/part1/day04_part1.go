package day04part1

import (
	"math"
	"reflect"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	"../../utils"
)

type date struct {
	Year  int
	Month int
	Day   int
	Hour  int
	Min   int
}

var wakesEvent = "wakes"
var fallsAsleepEvent = "asleep"
var beginShiftEvent = "#"

func guardStrategy1(records []string) int {
	defer utils.TimeTaken(time.Now())

	var sortedDates = make([]date, len(records))
	var events = make(map[date]string)
	var guards = make(map[string]map[int]int)
	var sleepTimeGuards = make(map[string]int)

	for i, line := range records {
		var date = getDate(line)
		sortedDates[i] = date
		var event = getEvent(line)
		events[date] = event
	}

	sort.SliceStable(sortedDates, func(i, j int) bool {
		return sortedDates[i].compareTo(sortedDates[j]) == -1
	})

	var currentGuardID string //= events[sortedDates[0]]
	var sleepStartTime = sortedDates[0]
	var sleepiestGuard = events[sleepStartTime]

	for _, currentDate := range sortedDates {
		if strings.ContainsAny(events[currentDate], beginShiftEvent) {
			currentGuardID = events[currentDate]
		} else {
			if _, knownGuard := guards[currentGuardID]; !knownGuard {
				guards[currentGuardID] = make(map[int]int)
			}

			if events[currentDate] == fallsAsleepEvent {

				if currentDate.Hour < 0 { //doesnt count before midnight
					currentDate.Hour = 0
					currentDate.Min = 0
				}

				sleepStartTime = currentDate
			} else if events[currentDate] == wakesEvent {
				sleepTimeGuards[currentGuardID] += currentDate.minutesSince(sleepStartTime)
				for minute := sleepStartTime.Min; minute < currentDate.Min; minute++ {
					guards[currentGuardID][minute]++
				}

				if sleepTimeGuards[currentGuardID] > sleepTimeGuards[sleepiestGuard] {
					sleepiestGuard = currentGuardID
				}
			}
		}
	}

	test := getKeyWithBiggerValue(guards[sleepiestGuard])
	test2, _ := strconv.Atoi(strings.Replace(sleepiestGuard, "#", "", -1))
	return test2 * test
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
