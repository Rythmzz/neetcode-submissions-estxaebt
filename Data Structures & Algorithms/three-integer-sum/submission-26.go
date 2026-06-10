func threeSum(nums []int) [][]int {

	sort.Ints(nums)
	result := [][]int{}

	for i:= 0;i < len(nums);i++ {
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}

		j,k := i+1, len(nums)-1

		for j < k {
			sum := nums[i] + nums[j] + nums[k]

			if sum < 0 {
				j++
			} else if sum > 0 {
				k--
			} else {
				for j < k && nums[j] == nums[j+1] {
					j++
				}

				for j < k && nums[k] == nums[k-1] {
					k--
				}

				arr := []int{nums[i],nums[j],nums[k]}
				result = append(result,arr)
				j++
				k--
			}
		}

	}

	return result

	
}
