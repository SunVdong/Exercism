package prime

import "errors"

// Nth returns the nth prime number. An error must be returned if the nth prime number can't be calculated ('n' is equal or less than zero)
func Nth(n int) (int, error) {
	if n <= 0 {
		return 0, errors.New("nth must be greater than zero")
	}

	count := 0
	for i := 2; ; i++ {
		if isPrime(i) {
			count++
		}
		if count == n {
			return i, nil
		}
	}
}

func isPrime(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
