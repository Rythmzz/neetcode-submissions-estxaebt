func findMin(nums []int) int {
	l,r := 0, len(nums)-1
	minNum := nums[0]

	for l <= r {
		mid := l + (r - l)/2
		fmt.Println(mid)
		if mid+1 < len(nums) && nums[mid] > nums[mid+1] {
			return nums[mid+1]
		} else if mid-1 >= 0 && mid+1 < len(nums) && nums[mid] < nums[mid+1] && nums[mid] < nums[mid-1]{
			return nums[mid]
		} else if minNum > nums[mid]{
			r--
		} else {
			l++
		}
	}

	return minNum
}
