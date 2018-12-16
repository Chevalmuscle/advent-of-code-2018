package day14

import (
	"strconv"
	"time"

	"../utils"
)

func recipes(input int) string {
	defer utils.TimeTaken(time.Now())

	var recipes = "37"
	var indexElf1 = 0
	var indexElf2 = 1

	for len(recipes) < input+10 {
		recipeElf1, _ := strconv.Atoi(string(recipes[indexElf1]))
		recipeElf2, _ := strconv.Atoi(string(recipes[indexElf2]))
		total := strconv.Itoa(recipeElf1 + recipeElf2)
		for _, newRecipe := range total {
			recipes += string(newRecipe)
		}
		indexElf1 = (indexElf1 + (recipeElf1 + 1)) % len(recipes)
		if indexElf1 == indexElf2 {
			indexElf1++
		}
		indexElf2 = (indexElf2 + (recipeElf2 + 1)) % len(recipes)
		if indexElf2 == indexElf1 {
			indexElf2++
		}
	}

	return recipes[input : input+10]
}
