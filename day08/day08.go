package day08

import (
	"strconv"
	"strings"
	"time"

	"../utils"
)

func memoryManeuver(input string) (int, int) {
	defer utils.TimeTaken(time.Now())

	var numbersString = strings.Split(input, " ")
	var numbers = make([]int, len(numbersString))
	for i, number := range numbersString {
		numbers[i], _ = strconv.Atoi(number)
	}

	sumPart1, sumPart2, _ := parseTree(numbers)

	return sumPart1, sumPart2
}

func parseTree(numbers []int) (int, int, int) {
	var numberOfChild = numbers[0]
	var numberOfMeta = numbers[1]
	var valuesOfChild = make(map[int]int)

	var startIndex = 2
	var rest = numbers[startIndex:]
	var sumPart1 = 0
	var sumPart2 = 0

	for i := 0; i < numberOfChild; i++ {
		sumTmp1, sumTmp2, indexTmp := parseTree(rest)
		valuesOfChild[i+1] = sumTmp2
		sumPart1 += sumTmp1
		startIndex += indexTmp
		rest = numbers[startIndex:]
	}

	for i := startIndex; i < startIndex+numberOfMeta; i++ {
		var metadata = numbers[i]
		if numberOfChild == 0 {
			sumPart2 += metadata
		} else {
			sumPart2 += valuesOfChild[metadata]
		}
		sumPart1 += metadata
	}

	return sumPart1, sumPart2, (startIndex + numberOfMeta)
}
