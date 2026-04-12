func twoSum(numbers []int, target int) []int {
	j := len(numbers) - 1
	i := 0

	for {
		if numbers[i]+numbers[j] == target {
			return []int{i+1, j+1}
		}
		if numbers[i]+numbers[j] > target {
			j--
		} else {
			i++
		}
	}
}
