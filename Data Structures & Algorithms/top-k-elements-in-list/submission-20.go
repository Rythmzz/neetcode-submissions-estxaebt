func topKFrequent(nums []int, k int) []int {
	freq := make(map[int]int)

	for _, n := range nums {
		freq[n]++
	}

	groups := make([][]int,len(nums)+1)

	for n, count := range freq {
		groups[count] = append(groups[count],n)
	}

	result := []int{}
	for i:= len(groups)-1;i >= 0;i-- {
		result = append(result,groups[i]...)
	}

	if len(result) > k {
		return result[:k]
	}
	
	return result
}
