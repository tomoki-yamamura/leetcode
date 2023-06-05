package dp

// coouldnt!!!!!!!
// func CanJump(nums []int) bool {
// 	dp := make(map[int]bool, len(nums)-1)
// 	for i := 0; i < len(nums); i++ {
// 		if !dp[i] && i != 0  {
// 			return false
// 		}
// 		for j := i+1; j <= nums[i]+i; j++ {
// 			if nums[j] == 0 {
// 				continue
// 			}
// 			dp[j] = true
// 		}
// 	}
// 	return true
// }

func CanJump(nums []int) bool {
	goal := len(nums) - 1
	for i := len(nums) - 1; i >= 0; i-- {
			if i+nums[i] >= goal {
					goal = i
			}
	}
	return goal == 0
}
