func largestRectangleArea(heights []int) int {
	n:= len(heights)
	left:= make([]int,n)
	right:= make([]int,n)
	maxAreaLarge := 0

	for i:= 0 ;i< n; i++{
		left[i] = -1
		right[i] = n
	}

	var stackR []int
	for i:= 0 ; i< n;i++ {

		for len(stackR) > 0 && heights[i] < heights[stackR[len(stackR)-1]] {
			idx := stackR[len(stackR)-1]
			stackR = stackR[:len(stackR)-1]

			right[idx] = i
		}
		stackR = append(stackR,i)
	}

	var stackL []int
	for i:= n-1 ; i >= 0;i-- {

		for len(stackL) > 0 && heights[i] < heights[stackL[len(stackL)-1]] {
			idx := stackL[len(stackL)-1]
			stackL = stackL[:len(stackL)-1]

			left[idx] = i
		}
		stackL = append(stackL,i)
	}

	for i:= 0 ;i< n;i++ {
		width := right[i] - left[i] - 1
		areaLarge := width * heights[i]
		maxAreaLarge  = max(maxAreaLarge,areaLarge)
	}

	return maxAreaLarge
}
