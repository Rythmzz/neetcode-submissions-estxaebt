func topKFrequent(nums []int, k int) []int {
	frequencyNums := map[int]int{}

    for _, num := range nums {
        frequencyNums[num]++
    }

    buckets := make([][]int, len(nums) + 1)
    result := make([]int,0,k)

    for num,count := range frequencyNums {
        buckets[count] = append(buckets[count],num)
    }

    for i:= len(buckets) - 1; i >= 0 && len(result) < k; i-- {
        result = append(result,buckets[i]...)
    }

    if len(result) > k {
        return result[:k]
    }

    return result
}
