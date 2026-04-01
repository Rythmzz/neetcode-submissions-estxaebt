func hasDuplicate(nums []int) bool {
    checker  := map[int]bool{}

	for _, num := range nums {
		if checker[num] {
			return true
		}
		checker[num] = true
	}

	return false

}
