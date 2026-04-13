
func maxArea(heights []int) int {
	left := 0
	right := len(heights) - 1
	max := -1
	for left < right {
		volume := (right - left) * min(heights[left], heights[right])
		if volume > max {
			max = volume
		}
		if heights[left] <= heights[right] {
			left++
		}else {
			right--
		}
	}
	return max
}
