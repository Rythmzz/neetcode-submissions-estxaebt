func longestConsecutive(nums []int) int {
	start := map[int]bool{}
	maxCount := 0

	for _, n:= range nums {
		start[n] = true
	}

	for _, n:= range nums {
		if !start[n-1] {
			count := 0
			for start[n+count] {
				count++
			}
			maxCount = max(maxCount,count)
		}
	}

	return maxCount
}
