package dp

// func NumDecodings(s string) int {
// 	dp := make(map[int]int)
// 	dp[len(s)] = 1

// 	var dfs func(int) int
// 	dfs = func(i int) int {
// 		if val, ok := dp[i]; ok {
// 			return val
// 		}
// 		if s[i] == '0' {
// 			return 0
// 		}

// 		res := dfs(i + 1)
// 		if i+1 < len(s) && (s[i] == '1' || (s[i] == '2' && s[i+1] >= '0' && s[i+1] <= '6')) {
// 			res += dfs(i + 2)
// 		}
// 		dp[i] = res
// 		return res
// 	}

// 	return dfs(0)
// }

// func NumDecodings(s string) int {
// 	if len(s) == 0 || s[0] == '0' {
// 			return 0
// 	}
// 	result := make([]int, len(s)+1)
// 	result[0] = 1
// 	for i := 1; i <= len(s); i++ {
// 			if s[i-1] != '0' {
// 					result[i] = result[i-1]
// 			}
// 			if i >= 2 && (s[i-2] == '1' || (s[i-2] == '2' && s[i-1] <= '6')) {
// 					result[i] += result[i-2]
// 			}
// 	}
// 	return result[len(s)]
// }

// func numDecodings(s string) int {
// 	if len(s) == 0 || s[0] == '0' {
// 		return 0
// 	}
// 	result := make([]int, len(s)+1)
// 	result[0] = 1

// }


func NumDecodings(s string) int {
	dp := make([]int, len(s)+1)
	dp[len(s)] = 1

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '0' {
			dp[i] = 0
		} else {
			dp[i] = dp[i+1]
			if i+1 < len(s) && (s[i] == '1' || (s[i] == '2' && s[i+1] >= '0' && s[i+1] <= '6')) {
				dp[i] += dp[i+2]
			}
		}
	}

	return dp[0]
}
