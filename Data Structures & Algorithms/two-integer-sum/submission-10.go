func twoSum(nums []int, target int) []int {
	checkNum := map[int]int{}

	for i := 0; i < len(nums); i++ {
		num := target - nums[i]
		idx, exists := checkNum[num]
		if exists {
			return []int{idx, i}
		}
		checkNum[nums[i]] = i
	}

	return nil
}
