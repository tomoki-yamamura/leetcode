package dp

// func rob(nums []int) int {
// 	if len(nums) == 0 { return 0 }
// 	if len(nums) == 1 { return nums[0] }

// 	return max(robWithoutCircle(nums[1:]), robWithoutCircle(nums[:len(nums)-1]))
// }

// func robWithoutCircle(nums []int) int {
// 	if len(nums) == 0 { return 0 }
// 	if len(nums) == 1 { return nums[0] }
// 	if len(nums) == 2 { return max(nums[0], nums[1]) }
	
// 	dp := make([]int, len(nums)+1)
// 	dp[0] = 0
// 	dp[1] = nums[0]
	
// 	for i := 1; i < len(nums); i++ {
// 			dp[i+1] = max(dp[i], nums[i] + dp[i-1])
// 	}
	
// 	return dp[len(nums)]
// }
