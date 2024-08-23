package spiralmatrix

func SpiralMatrix(size int) [][]int {
	// 初始化一个 size * size 的矩阵
	matrix := make([][]int, size)
	for i := range matrix {
		matrix[i] = make([]int, size)
	}

	row, col := 0, 0
	dr, dc := 0, 1
	for num := 1; num <= size*size; num++ {
		matrix[row][col] = num
		if row+dr < 0 || size <= row+dr || col+dc < 0 ||
			size <= col+dc || matrix[row+dr][col+dc] != 0 {
			dr, dc = dc, -dr
		}
		row, col = row+dr, col+dc
	}

	return matrix
}
