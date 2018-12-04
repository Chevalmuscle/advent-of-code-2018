package templatepart2

import (
	"strconv"
	"time"

	"../../utils"
)

func sum(numbers []string) int {
	defer utils.TimeTaken(time.Now())

	a, _ := strconv.Atoi(numbers[0])
	b, _ := strconv.Atoi(numbers[1])

	return a + b
}
