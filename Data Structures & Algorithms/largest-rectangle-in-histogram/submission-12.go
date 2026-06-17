func largestRectangleArea(heights []int) int {
	sl := make([]int,len(heights))
	sr := make([]int,len(heights))

	for i:= 0;i< len(heights);i++ {
		sl[i] = -1
		sr[i] = len(heights)
	}

	st := []int{}
	for i:= 0;i<len(heights);i++ {
		for len(st) > 0 && heights[i] < heights[st[len(st)-1]] {
			idx := st[len(st)-1]
			st = st[:len(st)-1]
			sr[idx] = i 
		}

		st = append(st,i)
	}

	st = []int{}
	for i:= len(heights)-1;i >= 0;i-- {
		for len(st) > 0 && heights[i] < heights[st[len(st)-1]] {
			idx := st[len(st)-1]
			st = st[:len(st)-1]
			sl[idx] = i 
		}

		st = append(st,i)
	}

	maxD := 0
	for i:= 0; i< len(heights);i++ {
		d := (sr[i] - sl[i] - 1) * heights[i]
		maxD = max(maxD,d)
	}

	return maxD
}
