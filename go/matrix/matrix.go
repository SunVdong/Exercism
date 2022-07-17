package matrix

import (
	"errors"
	"strconv"
	"strings"
)

// Define the Matrix type here.
type Matrix struct {
	rowNum int
	colNum int
	rows   [][]int
	cols   [][]int
}

func New(s string) (*Matrix, error) {
	rows := strings.Split(s, "\n")
	if len(rows) <= 0 {
		return nil, errors.New("wrong input")
	}
	m := Matrix{}
	m.rowNum = len(rows)
	for _, row := range rows {
		row = strings.Trim(row, " ")
		cols := strings.Split(row, " ")
		if m.colNum == 0 {
			m.colNum = len(cols)
		} else if m.colNum != len(cols) {
			return nil, errors.New("wrong")
		}
		var colArr []int
		for _, col := range cols {
			coli, e := strconv.Atoi(col)
			if e != nil {
				return nil, e
			}
			colArr = append(colArr, coli)
		}
		m.rows = append(m.rows, colArr)
	}

	return &m, nil
}

// Cols and Rows must return the results without affecting the matrix.
func (m *Matrix) Cols() [][]int {
	return nil
}

func (m *Matrix) Rows() [][]int {
	var res [][]int
	for _, v := range m.rows {
		res = append(res, v)
	}
	return res
}

func (m *Matrix) Set(row, col, val int) bool {
	m.rows[row-1][col-1] = val
	return true
}
