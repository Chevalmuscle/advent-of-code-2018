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
	var current = list.Front()
	var currentPlayer = 1

	//var winningPlayer = -1
	var winningScore = -1

	for marble := 1; marble <= lastMarble; marble++ {
		currentPlayer++
		if currentPlayer > playerCount {
			currentPlayer = 1
		}

		if marble%23 != 0 {
			var e = current.Next()
			if e == nil {
				e = list.Front()
			}
			current = list.InsertAfter(marble, e)
		} else {
			for i := 0; i < 7; i++ {
				current = current.Prev()
				if current == nil {
					current = list.Back()
				}
			}
			scores[currentPlayer] += current.Value.(int)
			scores[currentPlayer] += marble

			if scores[currentPlayer] > winningScore {
				winningScore = scores[currentPlayer]
				//winningPlayer = currentPlayer
			}
			var toRemove = current
			current = toRemove.Next()
			list.Remove(toRemove)

		}
	}
	/*
		for e := list.Front(); e != nil; e = e.Next() {
			fmt.Print(e.Value, " ")
		}*/

	//fmt.Println("\nwinning elf", winningPlayer)
	return winningScore
}
