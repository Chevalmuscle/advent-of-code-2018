package day08

import (
	"strconv"
	"strings"
	"time"

	"../utils"
)

func sumMetadata(input string) int {
	defer utils.TimeTaken(time.Now())

	var sum = 0

	var numbersString = strings.Split(input, " ")
	var numbers = make([]int, len(numbersString))
	for i, number := range numbersString {
		numbers[i], _ = strconv.Atoi(number)
	}

	sum, _ = recursive(numbers)
	return sum
}

func recursive(numbers []int) (int, int) {
	var numberOfChild = numbers[0]
	var numberOfMeta = numbers[1]

	var rest = numbers[2:]
	var sum = 0
	var index = 2
	for i := 0; i < numberOfChild; i++ {
		sumTmp, indexTmp := recursive(rest)
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
