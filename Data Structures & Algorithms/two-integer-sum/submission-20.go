func twoSum(nums []int, target int) []int {
    dup := map[int]int{} 

	for idx,num := range nums {
		dup[num] = idx+1
	}

	for idx,num := range nums {
		sum := target - num

		if dup[sum] != 0 && idx != dup[sum]-1 {
			return []int{idx,dup[sum]-1}
		}
	}

	return []int{}
}
