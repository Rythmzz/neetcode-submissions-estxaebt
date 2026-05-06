func productExceptSelf(nums []int) []int {
	mul := 1
	result := make([]int,len(nums))

	for i:= 0;i< len(nums);i++ {
		result[i] = mul
		mul *= nums[i]
	}

	mul = 1
	for i:= len(nums)-1; i >= 0;i-- {
		result[i] *= mul
		mul *= nums[i]
	}

	return result

}
