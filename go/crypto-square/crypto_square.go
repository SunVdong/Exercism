package cryptosquare

import (
	"strings"
)

func Encode(pt string) string {
	str := strings.Map(normalizer, pt)
	len := len(str)
	if len == 0 {
		return ""
	}

	r, c := getRowCol(len)

	var sb strings.Builder
	for col := 0; col < c; col++ {
		if col != 0 {
			sb.WriteString(" ")
		}
		for row := 0; row < r; row++ {
			idx := c*row + col
			if idx < len {
				r := str[idx]
				sb.WriteByte(r)
			} else {
				sb.WriteString(" ")
			}
		}
	}

	return sb.String()
}

func normalizer(r rune) (newR rune) {
	r, newR = r|32, -1
	if 'a' <= r && r <= 'z' || '0' <= r && r <= '9' {
		newR = r
	}
	return
}

func getRowCol(len int) (row, col int) {
	row = 1
	col = 1
	for {
		if row*col >= len {
			return
		}
		if col >= row+1 {
			row++
		} else {
			col++
		}
	}
}
