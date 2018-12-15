package day13

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"../utils"
)

const (
	DOWN       = "v"
	LEFT       = "<"
	UP         = "^"
	RIGHT      = ">"
	GOLEFT     = "left"
	GORIGHT    = "right"
	GOSTRAIGHT = "straight"
)

type point struct{ x, y int }
type value struct {
	value    string
	hasACart bool
}

type valueCart struct {
	value         string
	nextDirection string
}

var tracks map[point]value
var carts map[point]valueCart

var biggestX int
var biggestY int

func firstCrash(inputGraph []string) string {
	defer utils.TimeTaken(time.Now())
	var firstCrashLocation string

	tracks = make(map[point]value)
	carts = make(map[point]valueCart)

	biggestY = len(inputGraph)
	biggestX = len(inputGraph[0])

	for y := 0; y < len(inputGraph); y++ {
		currentLine := inputGraph[y]
		for x := 0; x < len(currentLine); x++ {

			currentPoint := point{x: x, y: y}
			currentCharacter := string(currentLine[x])
			if isCart(currentCharacter) {
				carts[currentPoint] = valueCart{value: currentCharacter, nextDirection: GOLEFT}
				if currentCharacter == LEFT || currentCharacter == RIGHT {
					tracks[currentPoint] = value{"-", true}
				} else {
					tracks[currentPoint] = value{"|", true}
				}
			} else if isIntersection(currentCharacter) {
				tracks[currentPoint] = value{currentCharacter, false}
			} else {
				tracks[currentPoint] = value{currentCharacter, false}
			}
		}
	}

	for firstCrashLocation == "" {
		newCarts := make(map[point]valueCart)

		for pos, cart := range carts {
			nextPos := nextPoint(pos, cart.value)
			valueOfPos := tracks[pos]
			valueOfPos.hasACart = false
			tracks[pos] = valueOfPos

			valueOfNextPos := tracks[nextPos]
			if valueOfNextPos.hasACart == true {
				firstCrashLocation = strconv.Itoa(nextPos.x) + "," + strconv.Itoa(nextPos.y)
			} else {
				valueOfNextPos.hasACart = true
			}
			if isIntersection(valueOfNextPos.value) {
				cart.value = turn(cart.value, cart.nextDirection)
				cart.iterateDirection()
				newCarts[nextPos] = cart
			} else if valueOfNextPos.value == "/" || valueOfNextPos.value == "\\" {
				turnDirection := getTurn(cart.value, valueOfNextPos.value)
				turnedCart := turn(cart.value, turnDirection)
				cart.value = turnedCart
				newCarts[nextPos] = cart
			} else {
				newCarts[nextPos] = cart
			}
			tracks[nextPos] = valueOfNextPos

		}
		carts = newCarts
		//printTracks()
	}

	return firstCrashLocation
}

func printTracks() {

	for y := 0; y < biggestY; y++ {
		for x := 0; x < biggestX; x++ {
			pos := point{x: x, y: y}
			if _, hasACart := carts[pos]; hasACart {
				myCart := carts[pos]
				fmt.Print(myCart.value)
			} else {
				fmt.Print(tracks[pos].value)
			}
		}
		fmt.Println("")
	}
}

func getTurn(cart string, track string) string {
	if cart == RIGHT && track == "\\" || cart == DOWN && track == "/" || cart == LEFT && track == "\\" || cart == UP && track == "/" {
		return GORIGHT
	} else if cart == RIGHT && track == "/" || cart == UP && track == "\\" || cart == LEFT && track == "/" || cart == DOWN && track == "\\" {
		return GOLEFT
	} else {
		log.Fatal("error in getTurn(...)")
		return ""
	}
}

func (v *valueCart) iterateDirection() {
	if v.nextDirection == GOLEFT {
		v.nextDirection = GOSTRAIGHT
	} else if v.nextDirection == GOSTRAIGHT {
		v.nextDirection = GORIGHT
	} else if v.nextDirection == GORIGHT {
		v.nextDirection = GOLEFT
	} else {
		log.Fatal("Next direction cannot be determined")
	}
}

func turn(cart string, direction string) string {
	if direction == GOLEFT {
		if cart == LEFT {
			return DOWN
		} else if cart == DOWN {
			return RIGHT
		} else if cart == RIGHT {
			return UP
		} else if cart == UP {
			return LEFT
		}
	} else if direction == GORIGHT {
		if cart == LEFT {
			return UP
		} else if cart == UP {
			return RIGHT
		} else if cart == RIGHT {
			return DOWN
		} else if cart == DOWN {
			return LEFT
		}
	}
	return cart
}

func nextPoint(currentPoint point, cart string) point {
	if cart == DOWN {
		return point{x: currentPoint.x, y: currentPoint.y + 1}
	} else if cart == LEFT {
		return point{x: currentPoint.x - 1, y: currentPoint.y}
	} else if cart == UP {
		return point{x: currentPoint.x, y: currentPoint.y - 1}
	} else if cart == RIGHT {
		return point{x: currentPoint.x + 1, y: currentPoint.y}
	} else {
		log.Fatal("Next point cannot be determined: not a cart")
		return point{}
	}
}

func isCart(str string) bool {
	return str == UP || str == DOWN || str == RIGHT || str == LEFT
}

func isIntersection(str string) bool {
	return str == "+"
}
