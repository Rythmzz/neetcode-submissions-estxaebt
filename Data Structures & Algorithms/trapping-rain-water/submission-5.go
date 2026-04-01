func trap(height []int) int {
	maxAmount := 0
	l := make([]int,len(height))
	r := make([]int,len(height))

	l[0] = height[0]
	for i:= 1 ; i < len(height);i++{
		l[i] = max(l[i-1],height[i])
	}

	r[len(height)-1] = height[len(height)-1]
	for i:= len(height)-2;i>= 0 ;i--{
		r[i] = max(r[i+1],height[i])
	}

	for i:= 0 ; i < len(height);i++{
		amount := min(r[i],l[i]) - height[i]
		maxAmount = maxAmount+amount
	}

	return maxAmount

}
