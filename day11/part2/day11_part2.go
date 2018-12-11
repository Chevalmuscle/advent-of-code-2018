package day11part2

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

	for y := 1; y <= maxSize; y++ {
		for x := 1; x <= maxSize; x++ {
			var sumSquare = getPowerLevel(x, y, serialNumber)

			if sumSquare > biggestPowerGrid {
				biggestPowerGrid = sumSquare
				corner = strconv.Itoa(x-1) + "," + strconv.Itoa(y) + ",1"
			}
			for squareSize := 2; squareSize < 20; squareSize++ {
				if x+squareSize > maxSize || y+squareSize > maxSize {
					break
				}
				sumSquare += getPowerAddedSize(x, y, serialNumber, squareSize)
				if sumSquare > biggestPowerGrid {
					biggestPowerGrid = sumSquare
					corner = strconv.Itoa(x) + "," + strconv.Itoa(y) + "," + strconv.Itoa(squareSize)
				}
			}
		}
	}
	return corner
}

func getPowerAddedSize(x int, y int, serialNumber int, squareSize int) int {
	var sum = -getPowerLevel(x+squareSize-1, y+squareSize-1, serialNumber)

	for i := y; i < y+squareSize; i++ {
		sum += getPowerLevel(x+squareSize-1, i, serialNumber)
	}
	for i := x; i < x+squareSize; i++ {
		sum += getPowerLevel(i, y+squareSize-1, serialNumber)
	}
	return sum
}

func getPowerLevel(x int, y int, serialNumber int) int {
	var rackID = x + 10
	var powerLevel = (((rackID * y) + serialNumber) * rackID) / 100
	if powerLevel > 9 {
		powerLevel %= 10
	}
	return powerLevel - 5
}
