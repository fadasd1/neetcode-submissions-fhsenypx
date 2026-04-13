func trap(height []int) int {
	i, j := 0, len(height) - 1
	iMax, jMax := height[0], height[len(height) - 1]
	water := 0

	for i < j {
		if height[i] < height[j] {
			if height[i] >= iMax {
				iMax = height[i]
			} else {
				water += iMax - height[i]
			}
			i++
		} else {
			if height[j] >= jMax {
				jMax = height[j]
			} else {
				water += jMax - height[j]
			}
			j--	
		}
		
	}
	return water
}