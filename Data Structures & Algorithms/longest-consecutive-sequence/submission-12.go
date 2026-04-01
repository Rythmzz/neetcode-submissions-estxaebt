func longestConsecutive(nums []int) int {
	dup := map[int]bool{}
	maxCount := 0

	for _, num:= range nums {
		dup[num] = true
	}

	for _, num := range nums {
		if !dup[num-1] {
			count := 1
			for dup[num+count] {
				count++
			}
			maxCount = max(maxCount,count)
		}
	}

	return maxCount
}
