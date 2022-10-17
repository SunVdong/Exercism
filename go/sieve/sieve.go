package sieve

func Sieve(limit int) []int {
	// default false; just ignores indexes 0, 1
	marks := make([]bool, limit+1)

	primes := make([]int, limit/2)
	pidx := 0
	for i := 2; i <= limit; i++ {
		if marks[i] {
			continue
		}
		primes[pidx] = i
		pidx++

		for m := i; m <= limit; m += i {
			// m is not prime, set to true
			marks[m] = true
		}
	}

	return primes[:pidx]
}
