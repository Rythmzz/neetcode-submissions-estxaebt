func maxArea(heights []int) int {
	l, r:= 0, len(heights)-1
	areaMax := 0
	for l < r {
		area := (r - l) * min(heights[l],heights[r])
		areaMax = max(areaMax,area)

		if heights[l] > heights[r] {
			r--
		} else if heights[r] > heights[l] {
			l++
		} else {
			l++
			r--
		}
	}

	return areaMax
}
