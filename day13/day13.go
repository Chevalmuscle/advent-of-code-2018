package day13

import (
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
type cartState struct{ value, nextDirection string }
type trackState struct {
	value    string
	hasACart bool
}

var tracks map[point]trackState
var carts map[point]cartState
var biggestX int
var biggestY int

func firstCrash(inputGraph []string) (string, string) {
	defer utils.TimeTaken(time.Now())

	var firstCrashLocation string
	tracks = make(map[point]trackState)
	carts = make(map[point]cartState)

	biggestY = len(inputGraph)
	biggestX = len(inputGraph[0])

	// parse the input
	for y := 0; y < len(inputGraph); y++ {
		currentLine := inputGraph[y]
		for x := 0; x < len(currentLine); x++ {

			currentPoint := point{x: x, y: y}
			currentCharacter := string(currentLine[x])
			if isCart(currentCharacter) {
				carts[currentPoint] = cartState{value: currentCharacter, nextDirection: GOLEFT}
				if currentCharacter == LEFT || currentCharacter == RIGHT {
					tracks[currentPoint] = trackState{"-", true}
				} else {
					tracks[currentPoint] = trackState{"|", true}
				}
			} else if isIntersection(currentCharacter) {
				tracks[currentPoint] = trackState{currentCharacter, false}
			} else {
				tracks[currentPoint] = trackState{currentCharacter, false}
			}
		}
	}

	// get these carts rollin
	for len(carts) > 1 {
		newCarts := make(map[point]cartState)

		for y := 0; y < biggestY; y++ {
			for x := 0; x < biggestX; x++ {
				pos := point{x: x, y: y}

				if _, cartIsOnTrack := carts[pos]; cartIsOnTrack {

					currentCartState := carts[pos]
					currentTrackValue := tracks[pos]
					currentTrackValue.hasACart = false
					tracks[pos] = currentTrackValue

					nextPos := nextPoint(pos, currentCartState.value)
					nextTrackValue := tracks[nextPos]

					if nextTrackValue.hasACart == true {
						// CRASH !
						if firstCrashLocation == "" {
							firstCrashLocation = strconv.Itoa(nextPos.x) + "," + strconv.Itoa(nextPos.y)
						}
						delete(newCarts, nextPos)
						delete(carts, nextPos)
						nextTrackValue.hasACart = false

					} else {
						nextTrackValue.hasACart = true
						if isIntersection(nextTrackValue.value) {
							// crosses an intersection
							currentCartState.value = turn(currentCartState.value, currentCartState.nextDirection)
							currentCartState.iterateDirection()
						} else if nextTrackValue.value == "/" || nextTrackValue.value == "\\" {
							// has to turn
							turnedCart := turn(currentCartState.value, getTurnDirection(currentCartState.value, nextTrackValue.value))
							currentCartState.value = turnedCart
						}
						newCarts[nextPos] = currentCartState

					}
					tracks[nextPos] = nextTrackValue
				}
			}
		}
		carts = newCarts
	}

	var lastCartLocation string
	for pos := range carts {
		lastCartLocation = strconv.Itoa(pos.x) + "," + strconv.Itoa(pos.y)
	}

	return firstCrashLocation, lastCartLocation
}

func getTurnDirection(cart string, track string) string {
	if cart == RIGHT && track == "\\" || cart == DOWN && track == "/" || cart == LEFT && track == "\\" || cart == UP && track == "/" {
		return GORIGHT
	} else if cart == RIGHT && track == "/" || cart == UP && track == "\\" || cart == LEFT && track == "/" || cart == DOWN && track == "\\" {
		return GOLEFT
	} else {
		log.Fatal("error in getTurnDirection(...)")
		return ""
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

func (v *cartState) iterateDirection() {
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
