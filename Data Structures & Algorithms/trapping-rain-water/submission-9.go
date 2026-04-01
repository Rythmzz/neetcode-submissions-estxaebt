func trap(height []int) int {
	lm := make([]int,len(height))
	rm := make([]int,len(height))
	countCol := 0

	lm[0] = height[0]
	for i:= 1 ;i< len(height);i++{
		lm[i] = max(height[i],lm[i-1])
	}

	rm[len(height)-1] = height[len(height)-1]
	for i:= len(height)-2 ;i >= 0;i--{
		rm[i] = max(height[i],rm[i+1])
	}

	for i:= 0; i< len(height);i++{
		col := min(rm[i],lm[i]) - height[i]
		countCol += col
	}

	return countCol
}
