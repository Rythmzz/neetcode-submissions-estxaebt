func twoSum(numbers []int, target int) []int {
		if len(numbers) == 0 {
		return nil
	}

	left, right := 0, len(numbers)-1
	for left < right {
		
		sum := numbers[left] + numbers[right]

		if sum > target {
			right--
		}

        if sum < target{
            left++
        }

		if sum == target {
			if left != right {
                return []int{left+1, right+1}
			}

		}

	}

	return nil



}
