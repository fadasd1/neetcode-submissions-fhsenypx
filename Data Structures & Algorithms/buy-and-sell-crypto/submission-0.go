func maxProfit(prices []int) int {
	buy := prices[0]

	currentBest := 0
	for i, val := range prices {
		sell := prices[i]
		profit := sell - buy
		if profit > currentBest {
			currentBest = profit
		}
		if val < buy {
			buy = val
		}
	}
	
	return currentBest
}
