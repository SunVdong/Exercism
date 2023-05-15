package series

func All(n int, s string) []string {
	l := len(s)
	if n > l {
		return []string{}
	}
	var res []string
	for i := 0; i <= l-n; i++ {
		res = append(res, s[i:i+n])
	}

	return res
}

func UnsafeFirst(n int, s string) string {
	if n>len(s) {
		panic(nil)
	}
	return s[0:n]
}
