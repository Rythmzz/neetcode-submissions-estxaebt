func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	minBuy := prices[0]
	totalPrice := 0

	for i:= 1 ;i < len(prices);i++ {
		profit := prices[i] - minBuy
		totalPrice = max(totalPrice,profit)

		minBuy = min(minBuy, prices[i])
	}

	return totalPrice
}
