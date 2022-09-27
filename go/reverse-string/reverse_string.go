package reverse

// Reverse returns the reversed string
func Reverse(s string) string {
	r := []rune(s)
	l := len(r)
	for i, j := 0, l-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
