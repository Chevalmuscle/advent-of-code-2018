package day06

import (
	"math"
	"strconv"
	"strings"
	"time"

	"../utils"
)

type point struct {
	x int
	y int
}

type pointValue struct {
	areaSize int
	isValid  bool
}

var biggestX = 0
var biggestY = 0

func sizeLargestArea(inputDangerousPoints []string) int {
	defer utils.TimeTaken(time.Now())

	var biggestAreaSize = 0
	var dangerousPoints = make(map[point]pointValue)

	for _, value := range inputDangerousPoints {
		var coordonates = strings.Split(value, ", ")
		var x, _ = strconv.Atoi(coordonates[0])
		var y, _ = strconv.Atoi(coordonates[1])
		dangerousPoints[point{x: x, y: y}] = pointValue{areaSize: 0, isValid: true}

		if x > biggestX {
			biggestX = x
		}
		if y > biggestY {
			biggestY = y
		}
	}

	for i := 0; i <= biggestX; i++ {
		for j := 0; j <= biggestY; j++ {
			var currentPoint = point{x: i, y: j}
			var closestpoint = point{x: math.MaxInt64, y: math.MaxInt64}
			var smallestDistance = math.MaxInt64
			var isEqual = false
			for dangerousPoint := range dangerousPoints {
				distance := manhattanDistance(currentPoint, dangerousPoint)

				if distance <= smallestDistance {
					if distance == smallestDistance {
						isEqual = true
					} else {
						isEqual = false
					}
					closestpoint = dangerousPoint
					smallestDistance = distance
				}
			}

			if !isEqual {
				// because golang...
				tmp := dangerousPoints[closestpoint]
				tmp.increaseAreaSize()
				dangerousPoints[closestpoint] = tmp

				if currentPoint.isBorder() {
					// because golang...
					tmp := dangerousPoints[closestpoint]
					tmp.setIsValid(false)
					dangerousPoints[closestpoint] = tmp
				}
			}
		}
	}

	for _, pointValue := range dangerousPoints {
		if pointValue.isValid {
			size := pointValue.areaSize
			if biggestAreaSize < size {
				biggestAreaSize = size
			}
		}
	}

	return biggestAreaSize

}

func (pv *pointValue) increaseAreaSize() {
	pv.areaSize++
}

func (pv *pointValue) setIsValid(isValid bool) {
	pv.isValid = isValid
}

func (p *point) isBorder() bool {
	return p.x == 0 || p.y == 0 || p.x == biggestX || p.y == biggestY
}
func manhattanDistance(p1 point, p2 point) int {
	return int(math.Abs(float64(p1.x-p2.x)) + math.Abs(float64(p1.y-p2.y)))
}
