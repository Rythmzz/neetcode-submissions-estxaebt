func maxProfit(prices []int) int {
	profitMax := 0
	buyMin := prices[0]
	
	for i:= 1;i< len(prices);i++ {
		profit := prices[i] - buyMin

		profitMax = max(profitMax,profit)
		buyMin = min(prices[i],buyMin)
	}

	return profitMax
}
