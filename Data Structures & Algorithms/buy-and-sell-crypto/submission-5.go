func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	profitMax := 0
	buyMin:= prices[0]

	for i:= 1;i < len(prices);i++{
		profit := prices[i] - buyMin
		profitMax = max(profitMax,profit)

		buyMin = min(buyMin,prices[i])
	}

	return profitMax

}
