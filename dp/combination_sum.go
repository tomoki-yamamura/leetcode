package dp

// //could not
// func combinationSum4(nums []int, target int) int {
// 	dp := make(map[int]int)
// 	dp[0] = 0
// 	for i := 1; i <= target; i++ {
// 		for j := 0 j <= i; j++ {
// 			dp[i] = dp[i-j]
// 		}
// 	}
// }

func CombinationSum4(nums []int, target int) int {
	dp := make(map[int]int)
	dp[0] = 1
	for i := 1; i <= target; i++ {
		dp[i] = 0
		for _, num := range nums {
			dp[i] += dp[i-num]
		}
	}
	return dp[target]
}