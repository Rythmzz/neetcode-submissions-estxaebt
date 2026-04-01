func trap(height []int) int {
	ml := make([]int,len(height))
	mr := make([]int,len(height))
	maxArea :=0

	ml[0] = height[0]
	for i:= 1 ; i< len(ml);i++ {
		ml[i] = max(height[i],ml[i-1])
	}

	mr[len(mr)-1] = height[len(height)-1]
	for i:= len(height)-2;i >= 0 ; i-- {
		mr[i] = max(height[i],mr[i+1])
	}

	for i:= 0 ; i< len(height);i++ {
		area := min(mr[i],ml[i])-height[i]
		maxArea += area
	}

	return maxArea
}
