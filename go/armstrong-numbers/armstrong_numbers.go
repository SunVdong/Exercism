package armstrong

import (
	"math"
	"strconv"
)

func IsNumber(n int) bool {
	power := float64(len(strconv.Itoa(n)))
	sum := 0
	x := n
	for x > 0 {
		sum += int(math.Pow(float64(x%10), power))
		x /= 10
	}
	return sum == n
}
