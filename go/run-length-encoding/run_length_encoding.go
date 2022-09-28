package encode

import (
	"strconv"
	"strings"
	"unicode"
)

func RunLengthEncode(s string) string {
	var encoded strings.Builder
	for len(s) > 0 {
		letter := s[0]
		slen := len(s)
		s = strings.TrimLeft(s, string(letter))
		if n := slen - len(s); n > 1 {
			encoded.WriteString(strconv.Itoa(n))
		}
		encoded.WriteByte(letter)
	}
	return encoded.String()
}

func RunLengthDecode(in string) string {
	multiplier := 0
	var out strings.Builder
	for _, r := range in {
		if unicode.IsDigit(r) {
			multiplier *= 10
			multiplier += int(r - '0')
		} else {
			if multiplier == 0 {
				multiplier++
			}
			out.WriteString(strings.Repeat(string(r), multiplier))
			multiplier = 0
		}
	}
	return out.String()
}

// Very clever!
// func RunLengthDecode(s string) string {
// 	var decoded strings.Builder
// 	for len(s) > 0 {
// 		i := strings.IndexFunc(s, func(r rune) bool {
// 			return !unicode.IsDigit(r)
// 		})
// 		n := 1
// 		if i != 0 {
// 			n, _ = strconv.Atoi(s[:i])
// 		}
// 		decoded.WriteString(strings.Repeat(string(s[i]), n))
// 		s = s[i+1:]
// 	}
// 	return decoded.String()
// }
