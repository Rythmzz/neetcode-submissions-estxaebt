func hasDuplicate(nums []int) bool {
    x := map[int]bool{}

	for i:= 0;i < len(nums);i++ {
		if !x[nums[i]] {
			x[nums[i]] = true
		} else {
			return true
		}
	}

	return false
}
