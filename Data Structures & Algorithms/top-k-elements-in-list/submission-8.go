func topKFrequent(nums []int, k int) []int {
    frequencyNum := map[int]int{}

    for _, num := range nums {
        frequencyNum[num]++
    }

    buckets := make([][]int,len(nums)+1)
    result := make([]int,0,k)

    for key, count := range frequencyNum {
        buckets[count] = append(buckets[count],key)
    }

    for i:= len(buckets) - 1; i >= 0; i--{
        result = append(result,buckets[i]...)
    }

    if len(result) > k {
        return result[:k]
    }

    return result
}
