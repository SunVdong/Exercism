package prime

import "errors"

var primes = []int{2}

func isPrime(n int) bool {
	for _, p := range primes {
		if n%p == 0 {
			return false
		}
	}
	return true
}

func Nth(n int) (int, error) {
	if n < 1 {
		return -1, errors.New("invalid input")
	}
	if n == 1 {
		return 2, nil
	}

	for i := 3; len(primes) < n; i++ {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}
	return primes[len(primes)-1], nil
}
