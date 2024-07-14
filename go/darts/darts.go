package darts

import "math"

func Score(x, y float64) int {
	des := math.Sqrt(x*x + y*y)
	if des <= 1 {
		return 10
	} else if des <= 5 {
		return 5
	} else if des <= 10 {
		return 1
	} else {
		return 0
	}
}
