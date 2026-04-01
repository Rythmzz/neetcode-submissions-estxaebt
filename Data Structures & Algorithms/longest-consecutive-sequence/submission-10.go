func longestConsecutive(nums []int) int {
	start := map[int]bool{}
	maxLength := 0

	for i:= 0 ;i < len(nums);i++ {
		start[nums[i]] = true
	}

	for i:= 0 ;i< len(nums);i++ {
		if !start[nums[i]-1] {
			count := 0
			for start[nums[i]+count] {
				count++
			}
			maxLength = max(maxLength,count)
		} 
	}

	return maxLength
}
