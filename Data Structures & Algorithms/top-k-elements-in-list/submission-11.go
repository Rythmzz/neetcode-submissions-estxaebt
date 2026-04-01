func topKFrequent(nums []int, k int) []int {
	frequent  := map[int]int{}

	for _, num:= range nums{
		frequent[num]++
	}

	groups := make([][]int,len(nums)+1)
	for num,count := range frequent {
		groups[count] = append(groups[count],num)
	}

	var result []int
	for i:= len(groups)-1;i>=0 && k > len(result);i--{
		result = append(result,groups[i]...)
	}

	if len(result) > k {
		return result[:k]
	}

	return result

}
