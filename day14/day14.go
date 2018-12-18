package day14

import (
	"log"
	"strconv"
	"strings"
	"time"

	"../utils"
)

func recipesPart1(input int) string {
	defer utils.TimeTaken(time.Now())

	var recipes = []int{3, 7}
	var indexElf1 = 0
	var indexElf2 = 1

	for len(recipes) < input+10 {
		recipeElf1 := recipes[indexElf1]
		recipeElf2 := recipes[indexElf2]

		newRecipies := getDigits(recipeElf1 + recipeElf2)
		recipes = append(recipes, newRecipies...)

		indexElf1 = (indexElf1 + (recipeElf1 + 1)) % len(recipes)
		indexElf2 = (indexElf2 + (recipeElf2 + 1)) % len(recipes)
	}

	var sb strings.Builder
	for i := 9; i >= 0; i-- {
		sb.WriteString(strconv.Itoa(recipes[len(recipes)-1-i]))
	}
	return sb.String()
}

func recipesPart2(input int) int {
	defer utils.TimeTaken(time.Now())

	var pattern = getDigits(input)
	var recipes = []int{3, 7}
	var indexElf1 = 0
	var indexElf2 = 1

	count := 0
	for count < 30000000 {
		recipeElf1 := recipes[indexElf1]
		recipeElf2 := recipes[indexElf2]

		newRecipies := getDigits(recipeElf1 + recipeElf2)
		recipes = append(recipes, newRecipies...)

		indexElf1 = (indexElf1 + (recipeElf1 + 1)) % len(recipes)
		indexElf2 = (indexElf2 + (recipeElf2 + 1)) % len(recipes)

		if len(pattern) < len(recipes) {
			if match(recipes, pattern) {
				return len(recipes) - len(pattern)
			} else if match(recipes[:len(recipes)-1], pattern) {
				return len(recipes) - len(pattern) - 1
			}
		}
		count++
	}

	// not found
	log.Fatalf("Pattern not found")
	return -1
}

func match(array1, array2 []int) bool {
	for i := 0; i < len(array2); i++ {
		teee := len(array1) - len(array2) + i
		if array1[teee] != array2[i] {
			return false
		}
	}
	return true
}

func getDigits(d int) []int {
	if d < 10 {
		return []int{d}
	}
	return append(getDigits(d/10), d%10)
}
