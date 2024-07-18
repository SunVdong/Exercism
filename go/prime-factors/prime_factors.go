package prime

func Factors(n int64) []int64 {
	var factors []int64
	for n%2 == 0 {
		factors = append(factors, 2)
		n /= 2
	}

	for i := int64(3); i*i <= n; i += 2 {
		for n%i == 0 {
			factors = append(factors, i)
			n /= i
		}
	}

	if n > 2 {
		factors = append(factors, n)
	}
	return factors
}
