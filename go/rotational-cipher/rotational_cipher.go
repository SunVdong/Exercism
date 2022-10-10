package rotationalcipher

import "strings"

func RotationalCipher(plain string, shiftKey int) string {
	var sb strings.Builder
	for _, r := range plain {
		if r >= 'a' && r <= 'z' {
			sb.WriteRune(rune((int(r-'a')+shiftKey)%26 + 'a'))
		} else if r >= 'A' && r <= 'Z' {
			sb.WriteRune(rune((int(r-'A')+shiftKey)%26 + 'A'))
		} else {
			sb.WriteRune(r)
		}
	}

	return sb.String()
}
