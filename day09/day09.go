package day09

import (
	"container/list"
	"regexp"
	"strconv"
	"time"

	"../utils"
)

func marbleScore(input string) int {
	defer utils.TimeTaken(time.Now())

	var inputs = regexp.MustCompile("(\\d+)").FindAllString(input, -1)
	var playerCount, _ = strconv.Atoi(inputs[0])
	var lastMarble, _ = strconv.Atoi(inputs[1])
	var scores = make(map[int]int)

	var list = list.New()
	list.PushBack(0)
	var currentMarble = list.Front()
	var currentPlayer = 1

	var winningScore = -1

	for marble := 1; marble <= lastMarble*100; marble++ { // remove *100 for part 1
		currentPlayer++
		if currentPlayer > playerCount {
			currentPlayer = 1
		}
		if marble%23 != 0 {
			var e = currentMarble.Next()
			if e == nil {
				e = list.Front()
			}
			currentMarble = list.InsertAfter(marble, e)
		} else {
			for i := 0; i < 7; i++ {
				currentMarble = currentMarble.Prev()
				if currentMarble == nil {
					currentMarble = list.Back()
				}
			}
			scores[currentPlayer] += currentMarble.Value.(int)
			scores[currentPlayer] += marble

			if scores[currentPlayer] > winningScore {
				winningScore = scores[currentPlayer]
			}
			var toRemove = currentMarble
			currentMarble = toRemove.Next()
			list.Remove(toRemove)
		}
	}

	return winningScore
}
