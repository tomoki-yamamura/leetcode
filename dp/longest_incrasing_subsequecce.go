package dp

import "sort"

// coudnt!!!!!!!
// func lengthOfLIS(nums []int) int {

// }

// o(n2)
// func LengthOfLIS(nums []int) int {
// 	dp := make([]int, len(nums))
//   for i := len(nums) - 1; i >= 0; i-- {
// 		dp[i] = 1
// 		for j := i+1; j < len(nums); j++ {
// 			if nums[i] < nums[j] {
// 				dp[i] = max(1 + dp[j], dp[i])
// 			}
// 		}
// 	}
// 	result := 0
// 	for _, num := range dp {
// 		result = max(num, result)
// 	}
// 	return result
// }

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

// o(nlogn)
func lengthOfLIS(nums []int) int {
	var piles [][]int
	for _, num := range nums {
			if len(piles) == 0 {
					piles = append(piles, []int{num})
					continue
			}
			j := sort.Search(len(piles), func(k int) bool {
					return piles[k][len(piles[k])-1] >= num
			})
			if j < len(piles) {
					piles[j] = append(piles[j], num)
			} else {
					piles = append(piles, []int{num})
			}
	}
	return len(piles)
}