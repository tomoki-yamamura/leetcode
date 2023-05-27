package dp

func Rob(nums []int) int {
  if len(nums) == 1 {
		return nums[0]
	} else if len(nums) == 2 {
		return max(nums[0], nums[1])
	}

	result := make([]int, len(nums))
	result[0] = nums[0]
	result[1] = max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		tmp := max(result[i-1], result[i-2] + nums[i])
		result[i] = tmp
	}
	return result[len(nums)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}