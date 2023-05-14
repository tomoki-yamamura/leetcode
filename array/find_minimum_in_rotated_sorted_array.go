package array


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

// solution
// func FindMin(nums []int) int {
// 	left, right := 0, len(nums)-1
// 	for left < right {
// 		mid := (left + right) / 2
// 		if nums[mid] > nums[right] {
// 			left = mid + 1
// 		} else {
// 			right = mid
// 		}
// 	}
// 	return nums[left]
// }
// arg := []int{3,4,5,6,7,8,1,2}
