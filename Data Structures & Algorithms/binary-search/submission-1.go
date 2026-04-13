func search(nums []int, target int) int {
	result := -1
	l := 0
	r := len(nums) - 1

	for l <= r  {
		middle := (r-l) / 2 + l
		if nums[middle] == target {
			return middle
		} else if nums[middle] > target {
			r = middle - 1
		} else if nums[middle] < target {
			l = middle + 1
		}
	}

	
	return result
}
