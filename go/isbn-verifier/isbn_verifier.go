package isbn

func IsValidISBN(isbn string) bool {
	count := 0
	sum := 0
	for _, v := range isbn {
		if (v >= '0' && v <= '9') || v == 'X' {
			count++
			var num int
			if v == 'X' {
				if count == 10 {
					num = 10
				} else {
					return false
				}
			} else {
				num = int(v) - int('0')
			}

			sum += num * (11 - count)
		} else if v != '-' {
			return false
		}
	}
	return count == 10 && sum%11 == 0
}
