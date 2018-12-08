package day06part2

import (
	"math"
	"strconv"
	"strings"
	"time"

	"../../utils"
)

type point struct {
	x int
	y int
}

var biggestX = 0
var biggestY = 0

func sizeLargestArea(inputSafePoints []string, maxSummedDistances int) int {
	defer utils.TimeTaken(time.Now())

	var areaSize = 0
	var safePoints = make(map[point]bool)

	for _, value := range inputSafePoints {
		var coordonates = strings.Split(value, ", ")
		var x, _ = strconv.Atoi(coordonates[0])
		var y, _ = strconv.Atoi(coordonates[1])
		safePoints[point{x: x, y: y}] = true

		if x > biggestX {
			biggestX = x
		}
		if y > biggestY {
			biggestY = y
		}
	}

	for i := 0; i <= biggestX; i++ {
		for j := 0; j <= biggestY; j++ {
			var sum = 0
			for safePoint := range safePoints {
				sum += manhattanDistance(point{x: i, y: j}, safePoint)
			}
			if sum < maxSummedDistances {
				areaSize++
			}
		}
	}

	return areaSize
}

func manhattanDistance(p1 point, p2 point) int {
	return int(math.Abs(float64(p1.x-p2.x)) + math.Abs(float64(p1.y-p2.y)))
}
