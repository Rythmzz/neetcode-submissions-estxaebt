func topKFrequent(nums []int, k int) []int {
	frequency := map[int]int{}

	for _, n:= range nums {
		frequency[n]++
	}

	groups := make([][]int,len(nums)+1)

	for num, count := range frequency {
		groups[count] = append(groups[count],num)
	}

	result := []int{}
	for i:= len(groups)-1;i>=0;i-- {
		result = append(result,groups[i]...)
	}

	if len(result) > k {
		return result[:k]
	}

	return result
}
