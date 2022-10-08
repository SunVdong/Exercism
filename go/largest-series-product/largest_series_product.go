package lsproduct

import (
	"errors"
)

func LargestSeriesProduct(digits string, span int) (int64, error) {
	if span < 0 {
		return 0, errors.New("invalid span")
	}
	if len(digits) < span {
		return 0, errors.New("span must be smaller than string length")
	}

	var max int64 = 0
	for i := 0; i <= len(digits)-span; i++ {
		subDigits := digits[i : i+span]
		var subProduct int64 = 1
		for _, d := range subDigits {
			if d < '0' || d > '9' {
				return 0, errors.New("invalid digits")
			}
			subProduct *= int64(d - '0')
		}
		if subProduct > max {
			max = subProduct
		}
	}

	return max, nil
}
