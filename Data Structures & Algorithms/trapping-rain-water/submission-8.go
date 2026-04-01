func trap(height []int) int {
	sl := make([]int,len(height))
	sr := make([]int,len(height))
	maxArea := 0

	sl[0] = height[0]
	for i:= 1 ;i < len(height);i++ {
		sl[i] = max(height[i],sl[i-1])
	}

	sr[len(height)-1] = height[len(height)-1]
	for i:= len(height)-2; i >= 0;i-- {
		sr[i] = max(height[i],sr[i+1])
	}

	for i:= 0 ; i< len(height);i++ {
		area := min(sr[i],sl[i]) - height[i]
		maxArea += area
	}

	return maxArea
}
