package day03part2

import (
	"regexp"
	"strconv"
	"time"

	"../../utils"
)

type point struct {
	x int
	y int
}

func getCleanClaimID(claims []string) int {
	defer utils.TimeTaken(time.Now())

	// points with a claim on it
	// key: point; value: map of the claim's IDs on that point
	claimedPoints := make(map[point]map[int]bool)

	// map of all the clean claims
	cleanClaims := make(map[int]bool)

	for _, line := range claims {
		re := regexp.MustCompile("([0-9]+)")

		currentID, _ := strconv.Atoi(re.FindAllString(line, -1)[0])
		startX, _ := strconv.Atoi(re.FindAllString(line, -1)[1])
		startY, _ := strconv.Atoi(re.FindAllString(line, -1)[2])
		width, _ := strconv.Atoi(re.FindAllString(line, -1)[3])
		height, _ := strconv.Atoi(re.FindAllString(line, -1)[4])

		cleanClaims[currentID] = true

		// iterate through the claim's area
		for i := startX; i < startX+width; i++ {
			for j := startY; j < startY+height; j++ {

				position := point{x: i, y: j}

				if _, alreadyClaimed := claimedPoints[position]; alreadyClaimed {

					// adds the current claim's ID on that point
					claimedPoints[position][currentID] = true

					// adds all the claims of that point in the overlapping claims's map
					for k := range claimedPoints[position] {
						delete(cleanClaims, k)
					}

				} else {
					// adds the claim's ID on the point
					claimedPoints[position] = map[int]bool{currentID: true}
				}
			}
		}
	}

	if len(cleanClaims) == 1 {
		// returns the only claim that is not overlapping
		for id := range cleanClaims {
			return id
		}
	}

	// no unique clean claim
	return -1

}
