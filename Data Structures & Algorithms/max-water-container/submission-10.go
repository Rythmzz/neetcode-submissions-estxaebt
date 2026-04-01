func maxArea(heights []int) int {
	l, r := 0, len(heights)-1
	maxAmount :=0

	for l < r {
		amount := (r - l) * min(heights[l],heights[r])
		maxAmount = max(maxAmount,amount)

		if l < r && heights[l] > heights[r] {
			r--
		} else if l < r && heights[l] < heights[r] {
			l++
		} else {
			l++
			r--
		}
	}

	return maxAmount
}
