func twoSum(numbers []int, target int) []int {
	l,r := 0, len(numbers)-1

	for l < r {
		sum := numbers[l] + numbers[r]

		if target == sum {
			return []int{l+1,r+1}
		}

		if target > sum {
			l++
		} else {
			r--
		}
	}

	return nil
}
