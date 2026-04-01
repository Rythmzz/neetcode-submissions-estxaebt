func trap(height []int) int {
	maxAreaWater := 0

	for i := 1; i < len(height)-1; i++ {

		left, right := i, i
		maxLeft, maxRight := height[i], height[i]

		for left >= 0 {
			maxLeft = max(maxLeft, height[left])
			left--
		}

		for right < len(height) {
			maxRight = max(maxRight, height[right])
			right++
		}

		maxAreaWater += min(maxLeft, maxRight) - height[i]
	}

	return maxAreaWater
}