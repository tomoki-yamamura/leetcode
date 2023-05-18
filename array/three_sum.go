package array

import "sort"

// couldnt
func ThreeSum(nums []int) [][]int {
	result := make([][]int, 0)
	if len(nums) < 3 {
		return result
	}

	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		leftPointer := i+1
		rightPointer := len(nums)-1
		for leftPointer < rightPointer {
			sum := nums[i] + nums[leftPointer] + nums[rightPointer]
			if sum == 0 {
				element := []int{nums[i], nums[leftPointer], nums[rightPointer]}
				result = append(result, element)
				rightPointer--
				leftPointer++
				for leftPointer < rightPointer && nums[rightPointer] == nums[rightPointer+1] {
					rightPointer--
				}
				for leftPointer < rightPointer && nums[leftPointer] == nums[leftPointer-1] {
					leftPointer++
				}
			} else if sum > 0 {
				rightPointer--
			} else if sum < 0 {
				leftPointer++
			}
		}
	}
	return result
}

// func threeSum(nums []int) [][]int {
// 	var results [][]int
// 	sort.Ints(nums)
// 	for i := 0; i < len(nums)-2; i++ {
// 		if i > 0 && nums[i] == nums[i-1] {
// 			continue//To prevent the repeat
// 		}
// 		target, left, right := -nums[i], i+1, len(nums)-1
// 		for left < right {
// 			sum := nums[left] + nums[right]
// 			if sum == target {
// 				results = append(results, []int{nums[i], nums[left], nums[right]})
// 				left++
// 				right--
// 				for left < right && nums[left] == nums[left-1] {
// 					left++
// 				}
// 				for left < right && nums[right] == nums[right+1] {
// 					right--
// 				}
// 			} else if sum > target {
// 				right--
// 			} else if sum < target {
// 				left++
// 			}
// 		}
// 	}
// 	return results
// }