package minesweeper

import (
	"strconv"
	"strings"
)

// Annotate returns an annotated board
func Annotate(board []string) []string {
	rows := len(board)
	if rows == 0 {
		return []string{}
	}
	cols := len(board[0])

	var res []string
	for row, line := range board {
		var s strings.Builder
		for col, val := range line {
			count := 0
			if val == ' ' {
				var rs, cs, re, ce int = 0, 0, rows - 1, cols - 1
				if row-1 > 0 {
					rs = row - 1
				}
				if col-1 > 0 {
					cs = col - 1
				}
				if row+1 < rows {
					re = row + 1
				}
				if col+1 < cols {
					ce = col + 1
				}
				for i := rs; i <= re; i++ {
					count += strings.Count(board[i][cs:ce+1], "*")
				}
				if count > 0 {
					s.WriteString(strconv.Itoa(count))
				} else {
					s.WriteRune(' ')
				}
			} else {
				s.WriteRune(val)
			}
		}
		res = append(res, s.String())
	}

	return res
}
