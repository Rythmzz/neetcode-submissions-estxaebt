func twoSum(nums []int, target int) []int {
    dup := map[int]int{}

	for i:= 0 ;i< len(nums);i++ {
		dup[nums[i]] = i + 1 
	}

	for i:= 0; i< len(nums);i++ {
		sum := target - nums[i] 

		if dup[sum] != 0 && dup[sum] - 1 != i {
			return []int{i,dup[sum]-1}
		}
	}

	return nil
}
