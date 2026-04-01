func topKFrequent(nums []int, k int) []int {
if len(nums) < k {
		return nil
	}

	result := []int{}
	countDuplicate := make(map[int]int)
	isAddDuplicate := make(map[int]bool)

	for i := 0; i < k; i++ {
		result = append(result, 0)
	}

	fmt.Println(result)

	for _, num := range nums {
		countDuplicate[num]++
	}

	for i := 0; i < k; i++ {
		for _, num := range nums {
			if countDuplicate[result[i]] <= countDuplicate[num] {
				if !isAddDuplicate[num] {
					result[i] = num
				}
			}
		}
		isAddDuplicate[result[i]] = true

	}

	return result
}
