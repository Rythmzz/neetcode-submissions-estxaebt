func twoSum(nums []int, target int) []int {
    matchSum := map[int]int{}

	for i:=0 ;i < len(nums);i++{
		matchSum[nums[i]] = i+1
	}

	for i:=0 ;i < len(nums);i++{
		num := target - nums[i]
		if matchSum[num] != 0 && matchSum[num]-1 !=  i {
			return []int{i,matchSum[num]-1}
		}
	}

	return nil
}
