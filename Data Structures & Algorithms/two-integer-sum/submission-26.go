func twoSum(nums []int, target int) []int {
    seen := make(map[int]int)

	for i,n := range nums {
		seen[n] = i+1
	}

	for i,n := range nums {
		num := target - n

		if seen[num] != 0 && seen[num]-1 != i {
			return []int{i,seen[num]-1}
		}
	}

	return nil
}
