package grains

import (
    "math"
    "errors"
)

func Square(number int) (uint64, error) {
    if number <= 0 || number > 64 {
		return 0, errors.New("number error")
	}
	var num uint64 = uint64(math.Pow(2,float64(number-1)))
	return num, nil
}

func Total() uint64 {
	return math.MaxUint64
}
