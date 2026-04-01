func hasDuplicate(nums []int) bool {
    dup := map[int]bool{}

	for _, num:= range nums{
		if dup[num] {
			return true
		}
		dup[num] = true
	}

	return false
}
