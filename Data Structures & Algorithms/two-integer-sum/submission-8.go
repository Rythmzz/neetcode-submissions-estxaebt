func twoSum(nums []int, target int) []int {
	checkNum := map[int]int{}
	for i := 0; i < len(nums); i++ {
		checkNum[nums[i]] = i + 1
	}

	for i := 0; i < len(nums); i++ {
		num := target - nums[i]
		if checkNum[num] != i+1 && checkNum[num] != 0 {
			return []int{i, checkNum[num] - 1}
		}
	}

	return nil
}
