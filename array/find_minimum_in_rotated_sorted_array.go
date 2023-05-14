package array

import "math"

// myRecursive solution
// func findMin(nums []int) int {
// 	if len(nums) == 2 {
// 		return int(math.Min(float64(nums[0]), float64(nums[1])))
// 	}
// 	middle := nums[len(nums)/2]
// 	left := nums[:len(nums)/2+1]
//   right := nums[len(nums)/2:]
// 	if middle >= left[0] && middle <= right[len(right)-1] {
// 		return nums[0]
// 	}
// 	if middle > right[len(right)-1] {
// 		return findMin(right)
// 	}
// 	return findMin(left)
// }
