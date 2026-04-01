func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result  := [][]int{}

	for i:= 0; i < len(nums);i++ {
		if i > 0 && nums[i] == nums[i-1]{
			continue
		}

		l,r := i+1, len(nums)-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]

			if sum < 0 {
				l++
			} else if sum > 0 {
				r--
			}  else {
				for l < r && nums[l] == nums[l+1] {
					l++
				}

				for l < r && nums[r] == nums[r-1] {
					r--
				}

				arr := []int{nums[i],nums[l],nums[r]}
				result = append(result,arr)

				l++
				r--
			}
		}
	}

	return result
}
