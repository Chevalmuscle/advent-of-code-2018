package day03

import (
	"regexp"
	"strconv"
	"time"

	"../utils"
)

type point struct {
	x int
	y int
}

func countOverlap(claims []string) int {
	defer utils.TimeTaken(time.Now())

	claimedPoints := make(map[point]bool)
	overlapAmount := 0

	for _, line := range claims {
		re := regexp.MustCompile("([0-9]+)")

		startX, _ := strconv.Atoi(re.FindAllString(line, -1)[1])
		startY, _ := strconv.Atoi(re.FindAllString(line, -1)[2])
		width, _ := strconv.Atoi(re.FindAllString(line, -1)[3])
		height, _ := strconv.Atoi(re.FindAllString(line, -1)[4])

		for i := startX; i < startX+width; i++ {
			for j := startY; j < startY+height; j++ {

				position := point{x: i, y: j}

				if _, claimed := claimedPoints[position]; claimed {

					// if the overlap was not already count
					if claimedPoints[position] {
						overlapAmount++
						claimedPoints[position] = false
					}

				} else {
					claimedPoints[position] = true
				}
			}
		}
	}
	return overlapAmount
}
