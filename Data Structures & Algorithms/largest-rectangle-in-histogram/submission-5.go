func largestRectangleArea(heights []int) int {
	sl := make([]int,len(heights))
	sr := make([]int,len(heights))

	for i:= 0 ; i< len(heights); i++ {
		sl[i] = -1
		sr[i] = len(heights)
	}

	stack := []int{}
	for i:= 0 ; i< len(heights);i++ {
		for len(stack) > 0 && heights[stack[len(stack)-1]] > heights[i] {
			idx := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			sr[idx] = i
		}
		stack = append(stack,i)
	}

	stack = []int{}
	for i:= len(heights)-1 ; i >= 0;i-- {
		for len(stack) > 0 && heights[stack[len(stack)-1]] > heights[i] {
			idx := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			sl[idx] = i
		}
		stack = append(stack,i)
	}

	maxRectangle := 0
	for i:= 0 ;i< len(heights);i++{
		rectangle :=  heights[i] * (sr[i] - sl[i] - 1)
		maxRectangle = max(maxRectangle, rectangle)
	}

	return maxRectangle

}
