package array

func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	tmp := nums[0]
	result := nums[0]
	for i := 1; i < len(nums); i++ {
		tmp = max(nums[i], tmp+nums[i]) // -1
		result = max(result, tmp) // 1
	}
	return tmp
}

func max(a, b int) int {
	if a > b {
		return a
	} 
	return b
}
