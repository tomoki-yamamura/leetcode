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





// func CombinationSum4(nums []int, target int) int {
// 	dp := make(map[int]int)
// 	dp[0] = 1
// 	for i := 1; i <= target; i++ {
// 		dp[i] = 0
// 		for _, num := range nums {
// 			if (i - num) >= 0 {
// 				dp[i] += dp[i-num]
// 			}
// 		}
// 	}
// 	return dp[target]
// }


// func CombinationSum4(nums []int, target int) int {
// 	dp := make(map[int]int)
// 	var dfs func(int)int
// 	dfs = func(want int) int {
// 		if want < 0 {
// 			return 0
// 		} else if want == 0 {
// 			return 1
// 		}

// 		if count, ok := dp[want]; ok {
// 			return count
// 		}

// 		count := 0
// 		for _, num := range nums {
// 			next := want - num
// 			count += dfs(next)
// 		}
// 		dp[want] = count
// 		return count
// 	}
// 	return dfs(target)
// }

func CombinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 1; i <= target; i++ {
		for _, num := range nums {
			if i-num >= 0 {
				dp[i] += dp[i-num]
			}
		}
	}
	return dp[target]
}
