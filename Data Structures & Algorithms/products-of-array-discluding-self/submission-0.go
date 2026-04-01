func productExceptSelf(nums []int) []int {
    var result []int

	index := 0
	for index < len(nums) {
		mul := 1
		for i := 0; i < len(nums); i++ {
			if index == i {
				continue
			} else {
				mul *= nums[i]
			}
		}
		result = append(result, mul)
		index++
	}

	return result
}
