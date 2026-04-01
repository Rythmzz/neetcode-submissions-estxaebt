func maxArea(heights []int) int {
	maxHeight := 0

	for left := 0; left < len(heights); left++ {
		right := len(heights) - 1
		for left < right {
			waterCal := (right-left) * min(heights[left],heights[right])
            maxHeight = max(maxHeight,waterCal)
			right--
		}
	}

	return maxHeight
}
