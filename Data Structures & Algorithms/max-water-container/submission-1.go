type Intermediary struct {
	i int
	j int
	maximum int
}

func maxArea(heights []int) int {
	left := 0
	right := len(heights) - 1
	max := Intermediary{
		i: -1,
		j: -1,
		maximum: 0,
	}
	for left < right {
		volume := (right - left) * min(heights[left], heights[right])
		if volume > max.maximum {
			max = Intermediary {
				i: left,
				j: right,
				maximum: volume,
			}
		}
		if heights[left] <= heights[right] {
			left++
		}else {
			right--
		}
	}
	return max.maximum
}
