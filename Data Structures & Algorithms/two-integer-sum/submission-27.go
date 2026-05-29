func twoSum(nums []int, target int) []int {
    seen := make(map[int]int)

	for i, n := range nums {
		seen[n] = i+1
	}

	for i, n := range nums {
		sum := target - n

		if seen[sum] != 0 && i != seen[sum]-1 {
			return []int{i,seen[sum]-1}
		}
	}

	return nil
}
