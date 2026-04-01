func maxProfit(prices []int) int {
	sellMin := prices[0]
	sellIndex := 0

	profitMax := 0

	for i:= 1 ; i < len(prices)-1;i++{
		if sellMin > prices[i]{
			sellMin = prices[i]
			sellIndex = i
		}
	}

	for i:= len(prices)-1 ; i > sellIndex;i--{
		profit := prices[i] - sellMin
		profitMax = max(profitMax,profit)
	}

	buyMax := prices[len(prices)-1]
	buyIndex := len(prices)-1

	for i:= len(prices)-1 ; i > 0;i--{
		if buyMax < prices[i]{
			buyMax = prices[i]
			buyIndex = i
		}
	}

	for i:= 0 ; i < buyIndex;i++{
		profit := buyMax - prices[i]
		profitMax = max(profitMax,profit)
	}


	return profitMax
}
