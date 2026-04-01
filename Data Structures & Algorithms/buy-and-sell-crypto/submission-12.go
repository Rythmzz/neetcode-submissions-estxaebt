func maxProfit(prices []int) int {
	minBuy := prices[0]
	profitMax := 0
	for i:= 1 ;i < len(prices);i++ {
		profit := prices[i] - minBuy
		profitMax = max(profitMax,profit)

		minBuy = min(minBuy,prices[i]) 
	}

	return profitMax
}
