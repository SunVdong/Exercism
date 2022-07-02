package luhn


func Valid(id string) bool {
	var sum, num int
	var count int = 0
	for i := len(id) - 1; i >= 0; i-- {
		c := id[i]
		switch {
		case c == ' ':
			continue
		case c >= '0' && c <= '9':
			num = int(c - '0')
			if count%2 == 1 {
				num <<= 1
			}
			if num > 9 {
				num -= 9
			}
			sum += num
			count++
		default:
			return false
		}
	}
	return count > 1 && sum%10 == 0
}
