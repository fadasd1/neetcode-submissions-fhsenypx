func twoSum(numbers []int, target int) []int {
	result := make([]int, 2)
	j := len(numbers) - 1
	i := 0

	for {
		if numbers[i]+numbers[j] == target {
			result[0], result[1] = i+1, j+1
			return result
		}
		if numbers[i]+numbers[j] > target {
			j--
		} else {
			i++
		}
	}
}
