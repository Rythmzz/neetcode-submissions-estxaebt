func longestConsecutive(nums []int) int {
	start := make(map[int]bool)

	for _, n := range nums {
		start[n] = true
	}

	length := 0
	for _, n := range nums {
		if !start[n-1] {
			count := 0
			for start[n+count] {
				count++
			}

			length = max(length,count)
		} 

	}

	return length
}
