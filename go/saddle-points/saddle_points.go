package matrix

import (
	"errors"
	"strconv"
	"strings"
)

// Define the Matrix and Pair types here.
type Matrix [][]int
type Pair [2]int

func New(s string) (*Matrix, error) {
	if s == "" {
		return &Matrix{}, nil
	}
	lines := strings.Split(s, "\n")
	if len(lines) <= 0 {
		return nil, errors.New("wrong")
	}

	res := make(Matrix, len(lines))
	for i, line := range lines {
		if line == "" {
			return nil, errors.New("wrong")
		}
		items := strings.Split(line, " ")
		if len(items) <= 0 {
			return nil, errors.New("wrong")
		}
		res[i] = make([]int, len(items))
		for j, item := range items {
			num, err := strconv.Atoi(item)
			if err != nil {
				return nil, errors.New("wrong")
			} else {
				res[i][j] = num
			}
		}
	}
	return &res, nil
}

func (m *Matrix) Saddle() []Pair {

	var res []Pair

	for i, line := range *m {
		maxInRow := line[0]
		var colIndex []int
		for j, val := range line {
			if val >= maxInRow {
				maxInRow = val
				colIndex = append(colIndex, j)
			}
		}

		for _, v := range colIndex {
			isSaddle := true
			for _, row := range *m {
				if row[v] < maxInRow {
					isSaddle = false
					break
				}
			}
			if isSaddle {
				res = append(res, Pair{i + 1, v + 1})
			}
		}

	}
	return res
}
