func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	l := 0
	r := len(matrix) - 1
	n := len(matrix[0]) - 1

	

	for l <= r {
		middle := l + (r-l) / 2
		
		left := matrix[middle][0]
		right := matrix[middle][n]
		if left <= target && right >= target {
			return search(matrix[middle], target)
		} else if left > target {
			r = middle - 1
		} else {
			l = middle + 1
		}

	}

	return false
}

func search(nums []int, target int) bool {
	l := 0
	r := len(nums) - 1

	for l <= r  {
		middle := (r-l) / 2 + l
		if nums[middle] == target {
			return true
		} else if nums[middle] > target {
			r = middle - 1
		} else if nums[middle] < target {
			l = middle + 1
		}
	}

	
	return false
}
