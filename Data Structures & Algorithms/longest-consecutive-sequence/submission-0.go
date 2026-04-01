func longestConsecutive(nums []int) int {
 	if len(nums) == 0 {
		return 0
	}
	sort.Ints(nums)

    // fmt.Println(nums)

	numExist := map[int]bool{
		nums[0]: true,
	}
	maxLength := 1
	count := nums[0]+1
	maxCount := 1

	for i := 1; i < len(nums); i++ {
		if numExist[nums[i]] {
			continue
		}

        // fmt.Printf("num[%d] = %d, count = % d\n",i,nums[i],count)

		if count == nums[i] {
			maxCount += 1
			numExist[nums[i]] = true
		} else {
			maxCount = 1
			count = nums[i]
		}

        maxLength = max(maxLength, maxCount)

		count++
	}

	return maxLength
}
