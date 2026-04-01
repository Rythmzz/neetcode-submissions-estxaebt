func trap(height []int) int {
	s:= make([]int,len(height))
	e:= make([]int,len(height))
	maxArea := 0

	s[0] = height[0]
	for i:= 1 ;i < len(height);i++{
		s[i] = max(s[i-1],height[i])
	}

	e[len(height)-1] = height[len(height)-1]
	for i:= len(height)-2; i >= 0 ;i-- {
		e[i] = max(e[i+1],height[i])
	}

	for i:= 0 ;i < len(height);i++ {
		area := min(e[i],s[i]) - height[i]
		maxArea += area
	}

	return maxArea
}
