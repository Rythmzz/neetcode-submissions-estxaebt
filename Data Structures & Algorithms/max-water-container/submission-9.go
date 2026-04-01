func maxArea(heights []int) int {
	maxAmount := 0

	l, r := 0, len(heights)-1
	for l < r {
		amount := (r - l) * min(heights[l],heights[r])
		maxAmount = max(maxAmount,amount)

		if heights[l] >  heights[r] {
			r--
		} else if heights[r] > heights[l] {
			l++
		} else {
			l++
			r--
		}


	}

	return maxAmount

}
