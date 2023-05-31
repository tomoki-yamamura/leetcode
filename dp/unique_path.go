package dp

func UniquePaths(m int, n int) int {
  grid := make([][]int, m)
	pre1 := make([]int, n)
	pre2 := make([]int, n)
	for i := 0; i < n; i++ {
		pre1[i] = 1
	}
	pre2[0] = 1
	for i := 0; i < m; i++ {
		if i == 0 {
			grid[i] = pre1
		} else {
			grid[i] = pre2
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			grid[i][j] = grid[i-1][j] + grid[i][j-1]
		}
	}

	return grid[m-1][n-1]
}
