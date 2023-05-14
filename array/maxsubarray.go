package array

import "math"

// couldnot!!!!!!!!!!!!!
// func maxProduct(nums []int) int {
//   result := nums[0]
// 	var tmp int
// 	var count int
// 	for i := 1; i < len(nums); i++ {
// 		if nums[i] == 0 {
// 			continue
// 		}
// 		if nums[i] < 0 {
// 			count++
// 			if count % 2 == 0 {
// 				result =
// 			}
// 			tmp = result * nums[i]
// 		}
// 		if nums[i] > 0 {
// 			result *= nums[i]
// 		}
// 	}
// 	return result
// }

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

// hint

func MaxProduct(nums []int) int {
	max := nums[0]
	min := nums[0]
	maxProduct := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			max, min = min, max
		}
		max = int(math.Max(float64(nums[i]), float64(max*nums[i])))
		min = int(math.Min(float64(nums[i]), float64(min*nums[i])))
		if max > maxProduct {
			maxProduct = max
		}
	}
	return maxProduct
}
