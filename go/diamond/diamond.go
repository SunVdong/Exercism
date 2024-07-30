package diamond

import (
	"bytes"
	"errors"
	"strings"
)

func Gen(char byte) (string, error) {
	if char < 'A' || char > 'Z' {
		return "", errors.New("wrong input")
	}

	num := int(char - 'A')
	colCount, rowCount := 2*num+1, 2*num+1
	rows := make([]string, rowCount)

	for c := byte('A'); c <= char; c++ {
		row := bytes.Repeat([]byte{' '}, colCount)
		row[char-c], row[colCount-1-int(char-c)] = c, c

		rows[c-'A'], rows[rowCount-1-int(c-'A')] = string(row), string(row)
	}

	return strings.Join(rows, "\n"), nil
}
