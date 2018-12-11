package day11part1

import (
	"math"
	"strconv"
	"time"

	"../../utils"
)

var maxSize = 300

func coordinates(serialNumber int) string {
	defer utils.TimeTaken(time.Now())

	var corner string
	var biggestPowerGrid = math.MinInt64

	for y := 1; y <= maxSize-3; y++ {
		var firstColumn = getPowerColumn(1, y, serialNumber)
		var secondColumn = getPowerColumn(2, y, serialNumber)
		var thirdColumm = getPowerColumn(3, y, serialNumber)

		for x := 1; x <= maxSize-3; x++ {
			var sum = firstColumn + secondColumn + thirdColumm

			if sum > biggestPowerGrid {
				biggestPowerGrid = sum
				corner = strconv.Itoa(x-1) + "," + strconv.Itoa(y)
			}

			firstColumn = secondColumn
			secondColumn = thirdColumm
			thirdColumm = getPowerColumn(x+2, y, serialNumber)
		}
	}
	return corner
}

func getPowerColumn(x int, y int, serialNumber int) int {
	var sum = 0
	for i := y; i < y+3; i++ {
		sum += getPowerLevel(x, i, serialNumber)
	}
	return sum
}

func getPowerLevel(x int, y int, serialNumber int) int {
	return (((x+10)*y)+serialNumber)*(x+10)/100%10 - 5
}
