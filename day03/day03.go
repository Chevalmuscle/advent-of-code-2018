package day03

import (
	"regexp"
	"strconv"
	"time"

	"../utils"
)

func countOverlap(claims []string) int {
	defer utils.TimeTaken(time.Now())

	claimedPositions := make(map[string]bool)
	overlapAmount := 0

	for _, line := range claims {
		re := regexp.MustCompile("([0-9]+)")

		startX, _ := strconv.Atoi(re.FindAllString(line, -1)[1])
		startY, _ := strconv.Atoi(re.FindAllString(line, -1)[2])
		width, _ := strconv.Atoi(re.FindAllString(line, -1)[3])
		height, _ := strconv.Atoi(re.FindAllString(line, -1)[4])

		for i := startX; i < startX+width; i++ {
			for j := startY; j < startY+height; j++ {

				//var position string
				position := strconv.Itoa(i) + "," + strconv.Itoa(j)

				if _, claimed := claimedPositions[position]; claimed {

					// if the overlap was not already count
					if claimedPositions[position] {
						overlapAmount++
						claimedPositions[position] = false
					}

				} else {
					claimedPositions[position] = true
				}
			}
		}
	}
	return overlapAmount
}
