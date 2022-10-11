package cryptosquare

import (
	"strings"
)

func Encode(pt string) string {
	str := normalize(pt)
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

func normalize(s string) string {
	var sb strings.Builder
	low := strings.ToLower(s)
	for _, r := range low {
		if r >= 'a' && r <= 'z' || r >= '0' && r <= '9' {
			// if unicode.IsLetter(r) || unicode.IsDigit(r) {
			sb.WriteRune(r)
		}
	}

	return sb.String()
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
