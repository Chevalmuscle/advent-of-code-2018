package day11

import (
	"math"
	"strconv"
	"time"

	"../utils"
)

func coordinates(serialNumber int) string {
	defer utils.TimeTaken(time.Now())

	var corner string
	var biggestPowerGrid int = math.MinInt64

	for y := 1; y <= 300-3; y++ {
		var firstColumn = getPowerColumn(1, y, serialNumber)
		var secondColumn = getPowerColumn(2, y, serialNumber)
		var thirdColumm = getPowerColumn(3, y, serialNumber)
		for x := 1; x <= 300-3; x++ {
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
	var rackID = x + 10
	var number = ((rackID * y) + serialNumber) * rackID
	var str = strconv.Itoa(number)
	var powerLevel int
	if len(str) > 2 {
		a, _ := strconv.Atoi(str[len(str)-3:])
		powerLevel = a / 100
	} else {
		powerLevel = 0
	}

	return powerLevel - 5
}
