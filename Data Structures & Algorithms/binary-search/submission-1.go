func search(nums []int, target int) int {
	mid := len(nums)/2
	end := len(nums)
	start := 0

	if target > nums[mid] {
		start = mid
	} else if target < nums[mid] {
		end = mid
	} else {
		return mid
	}

	for start < end {
		if target == nums[start] {
			return start
		}
		start++
	}

	return -1

}
