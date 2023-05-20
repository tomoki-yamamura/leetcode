package array

import "math"


func MaxArea(height []int) int {
  var result int
	left := 0;
	right := len(height)-1
	for left < right {
		horizontal := right - left
		vertical := math.Min(float64(height[left]), float64(height[right]))
		area := horizontal*int(vertical)
		result = int(math.Max(float64(area), float64(result)))
		if height[left] > height[right] {
			right--
			} else {
			left++
		}
	}
	return result
}

