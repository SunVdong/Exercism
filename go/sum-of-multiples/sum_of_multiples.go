package summultiples

func SumMultiples(limit int, divisors ...int) (sum int) {
	bits := make([]byte, (limit>>3)+1)

	for _, divisor := range divisors {
		if divisor == 0 {
			continue
		}
		for i := 1; ; i++ {
			num := divisor * i
			if num >= limit {
				break
			}

			index := num >> 3
			pos := num & 0x07

			if bits[index]&(1<<pos) == 0 {
				sum += num
				bits[index] |= 1 << pos
			}
		}
	}

	return
}
