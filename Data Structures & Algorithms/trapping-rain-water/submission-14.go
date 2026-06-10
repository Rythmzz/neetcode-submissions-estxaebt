func trap(height []int) int {
	hl,hr := make([]int,len(height)),make([]int,len(height))

	hl[0] = height[0]
	for i:= 1; i< len(height);i++ {
		hl[i] = max(hl[i-1],height[i])
	}

	hr[len(height)-1] = height[len(height)-1]
	for i:= len(height)-2;i >= 0;i-- {
		hr[i] = max(hr[i+1],height[i])
	}

	total := 0
	for i:= 0 ; i<len(height);i++ {
		area := min(hr[i],hl[i]) - height[i]
		total += area
	}

	return total
}
