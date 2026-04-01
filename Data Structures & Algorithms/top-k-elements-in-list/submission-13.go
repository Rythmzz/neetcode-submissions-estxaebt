func topKFrequent(nums []int, k int) []int {
	frequency := map[int]int{}

	for _ ,num := range nums {
		frequency[num]++
	}

	groups := make([][]int,len(nums)+1)
	for num,count := range frequency {
		groups[count] = append(groups[count],num)
	}

	var result  []int
	for i:= len(groups)-1;i >= 0 && len(result) < k ;i-- {
		result = append(result,groups[i]...)
	}

	if  len(result) > k {
		return result[:k]
	}

	return result

}
