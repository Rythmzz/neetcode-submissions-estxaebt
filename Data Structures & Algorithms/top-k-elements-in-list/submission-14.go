func topKFrequent(nums []int, k int) []int {
	freq := map[int]int{}

	for _, num:= range nums {
		freq[num]++
	}

	groups := make([][]int,len(nums)+1)
	result := []int{}

	for num, count := range freq {
		groups[count] = append(groups[count],num)
	}

	for i:= len(groups)-1; i >= 0 && len(result) < k; i--{
		result = append(result,groups[i]...)
	}

	if len(result) > k {
		return result[:k]
	}

	return result
}
