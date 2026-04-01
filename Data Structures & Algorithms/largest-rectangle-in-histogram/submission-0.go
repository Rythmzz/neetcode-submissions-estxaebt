func largestRectangleArea(heights []int) int {

	
	largest := 0

	for i := 0; i < len(heights); i++ {
		left, right := i, i
		for left >= 0 && heights[left] >= heights[i] {
			left--
		}

		for right < len(heights) && heights[right] >= heights[i] {
			right++
		}

		width := (right - left - 1) * heights[i]
		largest = max(largest, width)
	}

	return largest


}
