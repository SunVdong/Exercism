package pythagorean

type Triplet [3]int

func Range(min, max int) []Triplet {
	var triplet []Triplet
	for a := min; a < max; a++ {
		for b := a + 1; b < max; b++ {
			for c := max; c > b; c-- {
				sum := a*a + b*b
				if sum > c*c {
					break
				}
				if sum == c*c {
					triplet = append(triplet, Triplet{a, b, c})
				}
			}
		}
	}

	return triplet
}

func Sum(p int) (triplet []Triplet) {
	for c := p - 3; c > 2; c-- {
		for b := p - c - 1; b > 1; b-- {
			a := p - c - b
			if c <= b || b <= a {
				continue
			}
			sum := a*a + b*b
			if sum == c*c {
				triplet = append(triplet, Triplet{a, b, c})
			}
		}
	}

	return
}
