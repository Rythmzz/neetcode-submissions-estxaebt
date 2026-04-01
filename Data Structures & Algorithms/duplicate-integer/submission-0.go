func hasDuplicate(nums []int) bool {
    checkNum := map[int]bool{}
	for _, num := range nums {
		if checkNum[num] {
			return true
		}
		checkNum[num] = true
	}
	return false
}
