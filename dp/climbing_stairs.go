package dp

// bub
// func climbStairs(n int) int {
// 	if n == 0 {
// 		return 0
// 	}
// 	if n == 1 {
// 		return 1
// 	}
// 	if n == 2 {
// 		return 2
// 	}
// 	before1 := 2
// 	before2 := 1
// 	result := 0
// 	for i := 3; i < n; i++ {
// 		result = before1 + before2
// 		before1 = result
// 		before2 = before1 <<-
// 	}
// 	return result
// }

// collect & fast
// func climbStairs(n int) int {
// 	if n == 0 {
// 			return 0
// 	}
// 	if n == 1 {
// 			return 1
// 	}
// 	if n == 2 {
// 			return 2
// 	}
// 	before1 := 2
// 	before2 := 1
// 	result := 0
// 	for i := 3; i <= n; i++ {
// 			result = before1 + before2
// 			before2 = before1
// 			before1 = result
// 	}
// 	return result
// }

// func climbStairs(n int) int {
// 	if n == 0 {
// 		return 0
// 	}
// 	if n == 1 {
// 		return 1
// 	}
// 	if n == 2 {
// 		return 2
// 	}
// 	dp := make(map[int]int)
// 	dp[1] = 1
// 	dp[2] = 2
// 	for i := 3; i <= n; i++ {
// 		dp[i] = dp[i-1] + dp[i-2]
// 	}
// 	return dp[n]
// }
