package dp



// coudnt!!!!!!!
func LongestCommonSubsequence(text1 string, text2 string) int {
	max := func (a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	grid := make([][]int, len(text1)+1)
	for i := 0; i <= len(text1); i++ {
		grid[i] = make([]int, len(text2)+1)
	}

	for i := len(text1)-1; i >= 0; i-- {
		for j := len(text2)-1; j >= 0; j-- {
			if text1[i] == text2[j] {
				grid[i][j] = 1 + grid[i+1][j+1]
			} else {
				grid[i][j] = max(grid[i+1][j], grid[i][j+1])
			}
		}
	}
	return grid[0][0]
}

