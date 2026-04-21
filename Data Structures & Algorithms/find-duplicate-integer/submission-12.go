func findDuplicate(nums []int) int {
    f,s := nums[0],nums[0]

	for {
		s = nums[s]
		f = nums[nums[f]]
		if f == s {
			break
		}
	}

	s2 := nums[0]

	for s != s2 {
		s = nums[s]
		s2 = nums[s2]
	}

	return s
}
