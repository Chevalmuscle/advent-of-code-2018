package main

import (
	"fmt"
	"image/color"
	"math"
	"os"
	"regexp"
	"strconv"
	"time"

	"../utils"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
)

type point struct{ x, y int }
type speed struct{ dx, dy int }

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
		var dx, _ = strconv.Atoi(data[2])
		var dy, _ = strconv.Atoi(data[3])

		if _, knownPoint := points[point{x: x, y: y}]; knownPoint {
			points[point{x: x, y: y}] = append(points[point{x: x, y: y}], speed{dx: dx, dy: dy})
		} else {
			points[point{x: x, y: y}] = []speed{speed{dx: dx, dy: dy}}
		}
	}

	for i := 0; i < 15000; i++ {
		// to spot where the convergence will be
		if i%10 == 0 {
			xys := getPlotter(points)
			fileName := "images/out" + strconv.Itoa(i) + ".png"
			plotData(fileName, xys)
		}
		// 10240 is the second where the message is visible
		if i == 10240 {
			// final answer
			drawConsole(points)
		}

		points = update(points)
	}
}

func update(points map[point][]speed) map[point][]speed {
	var newPoints = make(map[point][]speed)

	smallestX = math.MaxInt64
	smallestY = math.MaxInt64
	biggestX = math.MinInt64
	biggestY = math.MinInt64

	for currentPoint, speeds := range points {
		for _, currentSpeed := range speeds {
			var newPoint = point{x: currentPoint.x + currentSpeed.dx, y: currentPoint.y + currentSpeed.dy}
			if _, knownPoint := newPoints[newPoint]; knownPoint {
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

func drawConsole(points map[point][]speed) {
	for y := smallestY; y <= biggestY; y++ {
		for x := smallestX; x <= biggestX; x++ {
			if _, isAPoint := points[point{x: x, y: y}]; isAPoint {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println("")
	}
}

func getPlotter(points map[point][]speed) plotter.XYs {
	var xys plotter.XYs

	for currentPoint := range points {
		xys = append(xys, struct{ X, Y float64 }{float64(currentPoint.x), float64(currentPoint.y)})
	}
	return xys
}

// comes from https://github.com/campoy/justforfunc/tree/master/34-gonum-plot
func plotData(path string, xys plotter.XYs) {
	f, _ := os.Create(path)
	p, _ := plot.New()
	s, _ := plotter.NewScatter(xys)
	s.Color = color.RGBA{R: 255, A: 255}
	p.Add(s)
	wt, _ := p.WriterTo(256, 256, "png")
	wt.WriteTo(f)
}
