func twoSum(nums []int, target int) []int {
    seen := map[int]int{}

	for i:= 0 ;i< len(nums);i++ {
		seen[nums[i]] = i + 1
	}

	for i:= 0 ; i< len(nums);i++ {
		num := target - nums[i]

		if seen[num] != 0 && seen[num]-1 != i {
			return []int{i,seen[num]-1}
		}
	}

	return nil
}
