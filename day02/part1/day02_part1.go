package main

import (
	"fmt"
	"time"

	"../../utils"
)

func main() {

	defer utils.TimeTaken(time.Now())

	occurenceTwoTime := 0
	occurenceThreeTime := 0

	lines := utils.ReadLines("../input.txt")

	for _, line := range lines {

		// key: rune; value: occurence of the rune
		var seenRunes = make(map[rune]int)

		for _, rune := range line {
			seenRunes[rune]++
		}

		// because it only counts once
		hasDoneTwoTime := false
		hasDoneThreeTime := false

		for _, value := range seenRunes {

			if value == 2 && !hasDoneTwoTime {
				hasDoneTwoTime = true
				occurenceTwoTime++
			} else if value == 3 && !hasDoneThreeTime {
				hasDoneThreeTime = true
				occurenceThreeTime++
			} else if hasDoneTwoTime && hasDoneThreeTime {
				// because it only counts once
				break
			}
		}
	}

	fmt.Printf("checksum: %d\n", occurenceTwoTime*occurenceThreeTime)

}
