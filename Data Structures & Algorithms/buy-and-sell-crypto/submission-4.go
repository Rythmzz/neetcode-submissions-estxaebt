func maxProfit(prices []int) int {
	sellMax := prices[0]
	sellIndex := 0

	profitMax := 0

	for i := 1; i < len(prices)-1; i++ {
		if sellMax > prices[i] {
			sellMax = prices[i]
			sellIndex = i
		}
	}

	for i := len(prices) - 1; i > sellIndex; i-- {
		profit := prices[i] - sellMax
		profitMax = max(profitMax, profit)
	}

	buyMin := prices[len(prices)-1]
	buyIndex := len(prices) - 1

	for i := len(prices) - 1; i > 0; i-- {
		if buyMin < prices[i] {
			buyMin = prices[i]
			buyIndex = i
		}
	}

	for i := 0; i < buyIndex; i++ {
		profit := buyMin - prices[i]
		profitMax = max(profitMax, profit)
	}

	return profitMax
}
