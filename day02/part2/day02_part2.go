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

	// Must choose between the two ways.
	// Where "n" is the amount of ids and "k" the length of one id.
	id1, id2 := hashMapWay(ids) // O(n^2 * k)
	//id1, id2 := naiveWay(ids) // O(n * k * log(k))

	fmt.Printf("IDs:\n\t%s \n\t%s\n\n", id1, id2)
	fmt.Printf("Common letters: \n\t%s\n\n", extractCommonLetters(id1, id2))

}

func hashMapWay(ids []string) (string, string) {
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
				return mapFirstHalf[id[:idLength/2]], id
			}
		}

		// the id is not in the map, so we add it
		mapFirstHalf[id[:idLength/2]] = id
	}

	// checks if the second half is the same
	for _, id := range ids {

		if _, isSameSecondtHalf := mapSecondHalf[id[idLength/2:]]; isSameSecondtHalf {
			// checks if the rest(first half) of the id has only one mistake
			if almostEqual(mapSecondHalf[id[idLength/2:]][:idLength/2], id[:idLength/2], mistakesAllowed) {
				return mapSecondHalf[id[idLength/2:]], id
			}
		}

		// the id is not in the map, so we add it
		mapSecondHalf[id[idLength/2:]] = id
	}

	// not found
	return "", ""
}

func naiveWay(ids []string) (string, string) {

	for i := 0; i < len(ids); i++ {
		for j := i; j < len(ids); j++ {
			if almostEqual(ids[i], ids[j], 1) {
				return ids[i], ids[j]
			}
		}
	}

	// not found
	return "", ""
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
