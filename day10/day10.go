package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"time"

	"../utils"
)

type point struct {
	x int
	y int
}

type speed struct {
	dx int
	dy int
}

var smallestX = math.MaxInt64
var smallestY = math.MaxInt64
var biggestX = math.MinInt64
var biggestY = math.MinInt64

func main() {
	defer utils.TimeTaken(time.Now())
	input := utils.ReadLines("input.txt")

	var points = make(map[point][]speed)

	var re = regexp.MustCompile("(-?\\d+)")
	for _, line := range input {
		data := re.FindAllString(line, -1)
		var x, _ = strconv.Atoi(data[0])
		var y, _ = strconv.Atoi(data[1])

		if x < smallestX {
			smallestX = x
		}
		if x > biggestX {
			biggestX = x
		}
		if y < smallestY {
			smallestY = y
		}
		if y > biggestY {
			biggestY = y
		}
		var dx, _ = strconv.Atoi(data[2])
		var dy, _ = strconv.Atoi(data[3])

		if _, isAPoint := points[point{x: x, y: y}]; isAPoint {
			points[point{x: x, y: y}] = append(points[point{x: x, y: y}], speed{dx: dx, dy: dy})
		} else {
			points[point{x: x, y: y}] = []speed{speed{dx: dx, dy: dy}}
		}

	}

	//f, _ := os.Create("day10_output")
	//defer f.Close()

	//w := bufio.NewWriter(f)
	for i := 0; i < 1; i++ {
		draw(points)
		//w.WriteString(output + "\n\n")
		//fmt.Printf("%s \n\n", output)
		points = update(points)
		//w.Flush()
	}

}

func update(points map[point][]speed) map[point][]speed {
	var newPoints = make(map[point][]speed)

	for currentPoint, speeds := range points {
		for _, currentSpeed := range speeds {
			var newPoint = point{x: currentPoint.x + currentSpeed.dx, y: currentPoint.y + currentSpeed.dy}
			if _, isAPoint := newPoints[newPoint]; isAPoint {
				newPoints[newPoint] = append(newPoints[newPoint], currentSpeed)
			} else {
				newPoints[newPoint] = []speed{currentSpeed}
			}

			if newPoint.x < smallestX {
				smallestX = newPoint.x
			}
			if newPoint.x > biggestX {
				biggestX = newPoint.x
			}
			if newPoint.y < smallestY {
				smallestY = newPoint.y
			}
			if newPoint.y > biggestY {
				biggestY = newPoint.y
			}
		}

	}

	return newPoints
}

func draw(points map[point][]speed) {
	//var sb strings.Builder

	for y := smallestY; y < biggestY; y++ {
		for x := smallestX; x < biggestX; x++ {
			if _, isAPoint := points[point{x: x, y: y}]; isAPoint {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println("")
	}

	//return sb.String()
}
