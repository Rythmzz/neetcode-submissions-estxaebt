func longestConsecutive(nums []int) int {
	checker := map[int]bool{}
	maxCount := 0

	for _, num:= range nums {
		checker[num] = true
	}

	for _, num:= range nums {
		if !checker[num-1] {
			count :=0
			for checker[num+count]{
				count++
			}
		maxCount =  max(maxCount,count)

		}
	}

	return maxCount
}
