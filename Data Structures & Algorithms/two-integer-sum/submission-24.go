func twoSum(nums []int, target int) []int {
    seen := map[int]int{}

	for i:= 0;i < len(nums);i++ {
		seen[nums[i]] = i + 1
	}

	for i:= 0 ;i< len(nums);i++ {
		sum := target - nums[i]
		if seen[sum] != 0 && (seen[sum]-1) != i {
			return []int{i,seen[sum]-1}
		}
	}

	return nil
}
