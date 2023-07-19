package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
	if n <= 0 {
		return 0, errors.New("Wrong input number")
	}

	if n == 1 {
		return 0, nil
	}

	count := 0
	for n != 1 {
		if n&1 == 0 {
			n >>= 1
			count++
		} else {
			n = n*3 + 1
			count++
		}
	}

	return count, nil
}
