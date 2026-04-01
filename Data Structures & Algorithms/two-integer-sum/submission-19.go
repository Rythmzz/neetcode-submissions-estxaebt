func twoSum(nums []int, target int) []int {
	checker := map[int]int{}

	for i:= 0; i <len(nums);i++{
		checker[nums[i]] = i+1
	}

	for i:= 0;i< len(nums);i++{
		result := target - nums[i]
		if checker[result] != 0 && i != checker[result]-1 {
			return []int{i,checker[result]-1}
		}
	}

	return nil
}
