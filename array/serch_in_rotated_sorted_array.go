package array

// couldnot!!!!!!!!!!!!!


// hint

// func search(nums []int, target int) int {
//   n := len(nums)
// 	left, right := 0, n-1
// 	for left > right {
// 		mid := left + (right-left)/2
// 		if nums[mid] > nums[right] {
// 			left = mid + 1
// 		} else {
// 			right = mid
// 		}
// 	}

// 	pivot := left
// }


// func search(nums []int, target int) int {
// 	var result int
// 	return helper(nums, target, result)
// }

// func helper(nums []int, target int, result int) int {
// 	if len(nums) == 1 {
// 		return result
// 	}
//   mid := len(nums)/2
// 	left := nums[:mid]
// 	right := nums[mid:]
// 	if left[0] <= target && target <= right[len(right)-1] {
// 		// mid = len(left)/2
// 		return helper(left, target, result)
// 	}
// 	result += mid
// 	return helper(right, target, result)
// }

// 4567012   

// func Search(nums[]int, target int) int {
// 	n := len(nums)
// 	left, right := 0, n-1
// 	for left < right { // left < right はpivotが一点になるまで
// 		mid := left+(right-left)/2
// 		if nums[mid] > nums[right] {
// 			left = mid + 1
// 		} else {
// 			right = mid
// 		}
// 	}

// 	pivot := left

// 	left, right = pivot, pivot+n-1
// 	for left <= right {
// 			mid := left+(right-left)/2
// 			midVal := nums[mid % n]
			
// 			if midVal > target {
// 					right = mid-1
// 			} else if midVal < target {
// 					left = mid+1
// 			} else {
// 					return mid % n
// 			}
// 	}
	
// 	return -1
// }

// arg := []int{4,5,6,7,0,1,2}
// arg := []int{4,5,6,7,0,1,2,4,5,6,7}
// result := array.Search(arg, 0)


// arg := []int{7,0,1,2,3,4,5,6} targ 3
func Search(nums []int, target int) int {    
	var lo, hi int = 0, len(nums) - 1
	
	for lo <= hi {
			var mid = lo + (hi - lo) / 2
			
			if nums[mid] == target {
					return mid
			}
			
			if nums[lo] <= nums[mid]{
					if target >= nums[lo] && target <= nums[mid] {
							hi = mid - 1
					} else {
							lo = mid + 1
					}
			}else if nums[mid] < nums[hi]{
					if target > nums[mid] && target <= nums[hi] {
							lo = mid + 1
					}else {
							hi = mid - 1
					}
			}
	}
	
	return -1
}