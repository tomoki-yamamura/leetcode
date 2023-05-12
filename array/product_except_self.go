package array


// couldnt
// func productExceptSelf(nums []int) []int {
  
// }

// example
func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums))
	pre := 1
	for i := 0; i < len(nums); i++ {
		result[i] = pre
		pre = pre * nums[i]
	}
	pre = 1
	for i := len(nums)-1; i >= 0; i-- {
		result[i] *= pre
		pre = pre * nums[i]
	}

	return result
}
