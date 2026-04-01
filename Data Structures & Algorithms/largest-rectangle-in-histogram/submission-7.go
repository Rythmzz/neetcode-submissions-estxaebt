func largestRectangleArea(heights []int) int {
	ls := make([]int,len(heights))
	rs := make([]int,len(heights))
	maxRectangle := 0

	for i:= 0 ;i< len(heights);i++ {
		ls[i] = -1
		rs[i] = len(heights)
	}

	stack := []int{}
	for i:= 0 ;i < len(heights);i++ {
		for len(stack) > 0 && heights[i] < heights[stack[len(stack)-1]]{
			idx := stack[len(stack)-1]
			rs[idx] = i
			stack = stack[:len(stack)-1]
		}
		stack = append(stack,i)
	}

	stack = []int{}
	for i:= len(heights)-1;i>=0;i-- {
		for len(stack) > 0 && heights[i] < heights[stack[len(stack)-1]]{
			idx := stack[len(stack)-1]
			ls[idx] = i
			stack = stack[:len(stack)-1]
		}
		stack = append(stack,i)
	}

	for i:= 0 ;i< len(heights);i++ {
		rectangle := heights[i] * (rs[i] - ls[i] - 1)
		maxRectangle = max(maxRectangle,rectangle)
	}

	return maxRectangle
}
