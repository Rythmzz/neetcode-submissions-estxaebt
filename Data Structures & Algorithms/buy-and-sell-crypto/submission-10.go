func maxProfit(prices []int) int {
	buyMin := prices[0]
	profitMax := 0
	for i:= 1 ; i< len(prices);i++ {
		profit := prices[i] - buyMin
		
		profitMax = max(profitMax,profit)
		buyMin = min(buyMin,prices[i])
	}

	return profitMax
}
