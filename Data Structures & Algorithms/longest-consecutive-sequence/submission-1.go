func longestConsecutive(nums []int) int {
	maxLength := 0
	set := map[int]bool{}

	for _, num := range nums {
		set[num] = true
	}

	var numSequence []int

	for _, num := range nums {
		if !set[num-1] {
			numSequence = append(numSequence, num)
		}
	}

	for _, num := range numSequence {
		num++
		count := 1
		for set[num] {
			count++
			num++
		}

		maxLength = max(maxLength, count)
	}

	return maxLength
}
