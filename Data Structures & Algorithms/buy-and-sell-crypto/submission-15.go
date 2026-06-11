func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	buyMin := prices[0]
	maxPrice := 0

	for _, price := range prices {
		maxPrice = max(maxPrice, price-buyMin)
		buyMin = min(buyMin,price)
	}

	return maxPrice

}
