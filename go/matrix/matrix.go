package matrix

import (
	"errors"
	"strconv"
	"strings"
)

// Define the Matrix type here.
type Matrix struct {
	data [][]int
}

func New(s string) (*Matrix, error) {
	rows := strings.Split(s, "\n")
	if len(rows) <= 0 {
		return nil, errors.New("wrong input")
	}
	m := Matrix{
		data: make([][]int, len(rows)),
	}
	var col_num int = 0
	for i, row := range rows {
		cols := strings.Split(strings.TrimSpace(row), " ")
		if col_num == 0 {
			col_num = len(cols)
		}
		if col_num != len(cols) {
			return nil, errors.New("all rows must have same width")
		}
		m.data[i] = make([]int, len(cols))
		for j, col := range cols {
			coli, e := strconv.Atoi(col)
			if e != nil {
				return nil, e
			}
			m.data[i][j] = coli
		}
	}

	return &m, nil
}

// Cols and Rows must return the results without affecting the matrix.
func (m *Matrix) Cols() [][]int {
	cols := make([][]int, len(m.data[0]))
	for i := range cols {
		cols[i] = make([]int, len(m.data))
		for j := range cols[i] {
			cols[i][j] = m.data[j][i]
		}
	}
	return cols
}

func (m *Matrix) Rows() [][]int {
	var r [][]int
	for _, row := range m.data {
		var rowdate []int
		rowdate = append(rowdate, row...)
		r = append(r, rowdate)
	}
	return r
}

func (m *Matrix) Set(row, col, val int) bool {
	row_size := len(m.data)
	col_size := len(m.data[0])
	if row < 0 || row >= row_size || col < 0 || col >= col_size {
		return false
	}
	m.data[row][col] = val
	return true
}
