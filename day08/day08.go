package day08

import (
	"strconv"
	"strings"
	"time"

	"../utils"
)

func sumMetadata(input string) (int, int) {
	defer utils.TimeTaken(time.Now())

	var numbersString = strings.Split(input, " ")
	var numbers = make([]int, len(numbersString))
	for i, number := range numbersString {
		numbers[i], _ = strconv.Atoi(number)
	}

	sumPart1, _ := recursivePart1(numbers)
	sumPart2, _ := recursivePart2(numbers)

	return sumPart1, sumPart2
}

func recursivePart2(numbers []int) (int, int) {
	var numberOfChild = numbers[0]
	var numberOfMeta = numbers[1]

	var valuesOfChild = make(map[int]int)
	var index = 2
	var rest = numbers[index:]
	var sum = 0

	for i := 0; i < numberOfChild; i++ {
		sumTmp, indexTmp := recursivePart2(rest)
		valuesOfChild[i+1] = sumTmp
		//sum += sumTmp
		index += indexTmp
		rest = numbers[index:]
	}

	var i = index
	for i < index+numberOfMeta {
		var metadata = numbers[i]
		//if _, isChild := valuesOfChild[metadata]; isChild {
		if numberOfChild == 0 {
			sum += metadata
		} else {
			sum += valuesOfChild[metadata]

		}
		//}
		i++
	}

	return sum, i
}

func recursivePart1(numbers []int) (int, int) {
	var numberOfChild = numbers[0]
	var numberOfMeta = numbers[1]
	var index = 2
	var rest = numbers[index:]
	var sum = 0
	for i := 0; i < numberOfChild; i++ {
		sumTmp, indexTmp := recursivePart1(rest)
		sum += sumTmp
		index += indexTmp
		rest = numbers[index:]
	}

	var i = index
	for i < index+numberOfMeta {
		sum += numbers[i]
		i++
	}

	return sum, i
}
