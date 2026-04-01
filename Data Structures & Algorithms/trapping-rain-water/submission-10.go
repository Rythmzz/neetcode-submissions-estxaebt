func trap(height []int) int {
	ml := make([]int,len(height))
	mr := make([]int,len(height))
	maxArea :=0

	ml[0] = height[0]
	for i:= 1;i < len(height);i++ {
		ml[i] = max(ml[i-1],height[i])
	}

	mr[len(height)-1] = height[len(height)-1]
	for i:= len(height)-2 ;i>= 0;i-- {
		mr[i] = max(mr[i+1],height[i])
	}

	for i:=0 ;i< len(height);i++{
		area := min(ml[i],mr[i]) - height[i]
		maxArea += area
	}

	return maxArea
}
