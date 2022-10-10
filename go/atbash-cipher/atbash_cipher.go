package atbash

import "strings"

var m = map[rune]rune{
	'a': 'z',
	'b': 'y',
	'c': 'x',
	'd': 'w',
	'e': 'v',
	'f': 'u',
	'g': 't',
	'h': 's',
	'i': 'r',
	'j': 'q',
	'k': 'p',
	'l': 'o',
	'm': 'n',
	'n': 'm',
	'o': 'l',
	'p': 'k',
	'q': 'j',
	'r': 'i',
	's': 'h',
	't': 'g',
	'u': 'f',
	'v': 'e',
	'w': 'd',
	'x': 'c',
	'y': 'b',
	'z': 'a',
	'0': '0',
	'1': '1',
	'2': '2',
	'3': '3',
	'4': '4',
	'5': '5',
	'6': '6',
	'7': '7',
	'8': '8',
	'9': '9',
}

func Atbash(s string) string {
	var sb strings.Builder
	lower := strings.ToLower(s)
	for _, r := range lower {
		if c, ok := m[r]; ok {
			if sb.Len()%6 == 5 {
				sb.WriteString(" ")
			}
			sb.WriteRune(c)
		}
	}

	return sb.String()
}
