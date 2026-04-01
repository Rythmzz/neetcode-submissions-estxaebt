func largestRectangleArea(heights []int) int {
	sl := make([]int,len(heights))
	sr := make([]int,len(heights))
	maxArea := 0

	for i:= 0 ;i< len(heights);i++ {
		sl[i] = -1
		sr[i] = len(heights)
	}

	stack := []int{}
	for i:= 0; i < len(heights);i++ {
		
		for len(stack) > 0 && heights[i] < heights[stack[len(stack)-1]] {
			idx := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			sr[idx] = i
		}
		stack = append(stack,i)
	}

	stack = []int{}
	for i:= len(heights)-1; i >= 0;i-- {
		
		for len(stack) > 0 && heights[i] < heights[stack[len(stack)-1]] {
			idx := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			sl[idx] = i
		}
		stack = append(stack,i)
	}

	for i:= 0 ; i < len(heights);i++ {
		area := heights[i] * (sr[i] - sl[i] - 1)
		maxArea = max(maxArea,area)
	}

	return maxArea
}
