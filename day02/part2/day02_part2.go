package main

import (
	"fmt"
	"strings"
	"time"

	"../../utils"
)

func main() {

	defer utils.TimeTaken(time.Now())

	ids := utils.ReadLines("../input.txt")
	mapFirstHalf := make(map[string]string)
	mapSecondHalf := make(map[string]string)

	mistakesAllowed := 1

	// length of one id
	idLength := len(ids[0])

	// checks if the first half is the same
	for _, id := range ids {

		if _, isSameFirstHalf := mapFirstHalf[id[:idLength/2]]; isSameFirstHalf {
			// checks if the rest(second half) of the id has only one mistake
			if almostEqual(mapFirstHalf[id[:idLength/2]][idLength/2:], id[idLength/2:], mistakesAllowed) {
				fmt.Printf("IDs:\n\t%s \n\t%s\n\n", mapFirstHalf[id[:idLength/2]], id)
				fmt.Printf("Common letters: \n\t%s\n\n", extractCommonLetters(mapFirstHalf[id[:idLength/2]], id))
				return
			}
		}

		// the id is not in the map, so we add it
		mapFirstHalf[id[:idLength/2]] = id
	}

	// checks if the second half is the same
	for _, id := range ids {

		if _, isSameSecondtHalf := mapSecondHalf[id[idLength/2+1:]]; isSameSecondtHalf {
			// checks if the rest(first half) of the id has only one mistake
			if almostEqual(mapSecondHalf[id[idLength/2:]][:idLength/2], id[:idLength/2], mistakesAllowed) {
				fmt.Printf("IDs:\n %s \n%s\n", mapSecondHalf[id[idLength/2:]], id)
				fmt.Printf("common letters: %s\n", extractCommonLetters(mapFirstHalf[id[:idLength/2]], id))
				return
			}
		}

		// the id is not in the map, so we add it
		mapSecondHalf[id[:idLength/2]] = id
	}
}

// extractSameLetters returns a string with the common
// letters of the two strings. The order is important.
// str 1 and str 2 must have the same length
func extractCommonLetters(str1 string, str2 string) string {

	var stringBuilder strings.Builder
	for i := 0; i < len(str1); i++ {
		if str1[i] == str2[i] {
			stringBuilder.WriteString(string(str1[i]))
		}
	}

	return stringBuilder.String()
}

// str 1 and str 2 must have the same length
func almostEqual(str1 string, str2 string, mistakesAllowed int) bool {

	for i := 0; i < len(str1); i++ {
		if str1[i] != str2[i] {
			mistakesAllowed--
		}
		if mistakesAllowed < 0 {
			break
		}
	}

	return mistakesAllowed == 0
}
