package allyourbase

import (
	"errors"
)

func ConvertToBase(inputBase int, inputDigits []int, outputBase int) ([]int, error) {
	if inputBase < 2 {
		return nil, errors.New("input base must be >= 2")
	}

	if outputBase < 2 {
		return nil, errors.New("output base must be >= 2")
	}

	sum := 0
	for _, val := range inputDigits {
		if val < 0 || val >= inputBase {
			return nil, errors.New("all digits must satisfy 0 <= d < input base")
		}
		// sum += val * int(math.Pow(float64(inputBase), float64(len-(idx+1))))
		sum = sum*inputBase + val
	}

	if sum == 0 {
		return []int{0}, nil
	}

	var res = []int{}
	for ; sum > 0; sum /= outputBase {
		res = append([]int{sum % outputBase}, res...)
	}

	return res, nil
}
