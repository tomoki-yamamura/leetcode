package array

func TwoSum(nums []int, target int) []int{
	box := make(map[int]int)
	for i, num := range nums {
		if ind, ok := box[target-num]; ok {
			return []int{ind, i}
		}
		box[num] = i
	}
	return []int{}
}