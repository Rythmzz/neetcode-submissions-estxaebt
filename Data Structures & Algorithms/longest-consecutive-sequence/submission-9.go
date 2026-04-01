func longestConsecutive(nums []int) int {
	start := map[int]bool{}
	maxCount := 0

	for _, num := range nums {
		start[num] = true
	}

	for _, num := range nums {
		if !start[num-1] {
			count := 0 
			for start[num+count] {
				count++
			}

			maxCount = max(maxCount,count)
		}
	}

	return maxCount
}
