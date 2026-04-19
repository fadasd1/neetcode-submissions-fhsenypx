
func minEatingSpeed(piles []int, h int) int {
    sum := 0
    max := piles[0]
    for _, val := range piles {
        if val > max {
            max = val
        }
		sum += val
	}
	left, right := ceiling(sum, h), max
	

	for left < right {
        rate := left + (right-left) / 2
        hours := 0
        for _, val := range piles {
            hours += ceiling(val, rate)            
        }
        if hours <= h {
                right = rate
            } else {
                left = rate + 1
            }
    }

	return left
}

func ceiling(a,b int) int {
    return (a + b - 1) / b
}

