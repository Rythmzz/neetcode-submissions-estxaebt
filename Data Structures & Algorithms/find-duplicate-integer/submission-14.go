func findDuplicate(nums []int) int {
    fast, slow := nums[0], nums[0]

	for {
		slow = nums[slow]
		fast = nums[nums[fast]]

		if slow == fast {
			break
		}
	}

	slow2 := nums[0]

	for slow != slow2 {
		slow = nums[slow]
		slow2 = nums[slow2]
	}

	return slow2
}
