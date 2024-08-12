package yacht

func Score(dice []int, category string) int {
	scoreFunc := func(num int) int {
		count := 0
		for _, v := range dice {
			if v == num {
				count++
			}
		}
		return num * count
	}

	isYachtFunc := func() bool {
		for _, v := range dice {
			if dice[0] != v {
				return false
			}
		}
		return true
	}

	countFunc := func() map[int]int {
		m := make(map[int]int)
		for _, v := range dice {
			m[v]++
		}
		return m
	}
	switch category {
	case "ones":
		return scoreFunc(1)
	case "twos":
		return scoreFunc(2)
	case "threes":
		return scoreFunc(3)
	case "fours":
		return scoreFunc(4)
	case "fives":
		return scoreFunc(5)
	case "sixes":
		return scoreFunc(6)
	case "full house":
		m := countFunc()
		var found_3, found_2 bool = false, false
		var key_3, key_2 = 0, 0
		for k, v := range m {
			if v == 3 {
				found_3, key_3 = true, k
			}
			if v == 2 {
				found_2, key_2 = true, k
			}
		}
		if found_2 && found_3 {
			return key_2*2 + key_3*3
		}
		return 0
	case "four of a kind":
		m := countFunc()
		found_4, key_4 := false, 0
		for k, v := range m {
			if v >= 4 {
				found_4, key_4 = true, k
			}
		}
		if found_4 {
			return 4 * key_4
		}
		return 0
	case "choice":
		sum := 0
		for _, v := range dice {
			sum += v
		}
		return sum
	case "little straight":
		m := countFunc()
		if m[1] != 1 {
			return 0
		}
		if m[2] != 1 {
			return 0
		}
		if m[3] != 1 {
			return 0
		}
		if m[4] != 1 {
			return 0
		}
		if m[5] != 1 {
			return 0
		}
		return 30
	case "big straight":
		m := countFunc()
		if m[2] != 1 {
			return 0
		}
		if m[3] != 1 {
			return 0
		}
		if m[4] != 1 {
			return 0
		}
		if m[5] != 1 {
			return 0
		}
		if m[6] != 1 {
			return 0
		}
		return 30
	case "yacht":
		if isYachtFunc() {
			return 50
		} else {
			return 0
		}
	default:
		return 0
	}

}
