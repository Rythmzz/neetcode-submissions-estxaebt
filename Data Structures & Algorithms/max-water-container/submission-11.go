func maxArea(heights []int) int {
	l, r := 0, len(heights)-1
	maxArea := 0

	for l < r {
		area := (r - l) * min(heights[l],heights[r])
		fmt.Printf("r %d, l%d, area = %d\n",r,l,area)
		maxArea = max(maxArea,area)

		if heights[l] > heights[r] {
			r--
		} else if heights[l] < heights[r] {
			l++
		} else {
			l++
			r--
		}
	}

	return maxArea
}
