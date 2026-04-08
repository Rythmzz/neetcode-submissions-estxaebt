func maxArea(heights []int) int {
	l,r := 0, len(heights)-1
	maxHeight := 0

	for l < r {
		height := (r-l) * min(heights[l],heights[r])
		maxHeight = max(maxHeight,height)

		if heights[l] > heights[r] {
			r--
		} else if heights[l] < heights[r] {
			l++
		} else {
			r--
			l++
		}
	}

	return maxHeight
}
