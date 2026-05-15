func largestRectangleArea(heights []int) int {
	lm := make([]int,len(heights))
	rm := make([]int,len(heights))

	for i:= 0;i< len(heights);i++ {
		lm[i] = -1
		rm[i] = len(heights)
	}

	st := []int{}
	for i:= 0; i < len(heights);i++ {
		for len(st) > 0 && heights[i] < heights[st[len(st)-1]] {
			idx := st[len(st)-1]
			st = st[:len(st)-1]
			rm[idx] = i
		}

		st = append(st,i)
	}

	st = []int{}
	for i:= len(heights)-1; i >= 0 ;i-- {
		for len(st) > 0 && heights[i] < heights[st[len(st)-1]] {
			idx := st[len(st)-1]
			st = st[:len(st)-1]
			lm[idx] = i
		}

		st = append(st,i)
	}

	result := 0
	for i:= 0; i < len(heights); i++ {
		rectangle := (rm[i] - lm[i] - 1) * heights[i]
		result = max(result,rectangle)
	}

	return result
}
