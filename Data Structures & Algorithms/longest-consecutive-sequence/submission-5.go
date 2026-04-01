func longestConsecutive(nums []int) int {
    seen := map[int]bool{}
    maxLength := 0

    for _, num := range nums {
        seen[num] = true
    }

    for _, num := range nums {
        if !seen[num-1] {
            count := 0
            for seen[num+count] {
                count++
            }
            maxLength = max(maxLength,count)
        }
    }

    return  maxLength
}
