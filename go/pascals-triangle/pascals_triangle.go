package pascal

func Triangle(n int) [][]int {
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		row := make([]int, i+1)
		row[0] = 1
		row[i] = 1
		for j := 1; j < i; j++ {
			row[j] = res[i-1][j-1] + res[i-1][j]
		}

		res[i] = row
	}

	return res
}
