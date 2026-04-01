func hasDuplicate(nums []int) bool {
	dup := map[int]bool{}

	for i:= 0 ;i < len(nums);i++ {
		if dup[nums[i]] {
			return true
		}

		dup[nums[i]] = true
	}

	return false
}
