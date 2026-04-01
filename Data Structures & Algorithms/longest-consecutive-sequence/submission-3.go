func longestConsecutive(nums []int) int {
    checkNum := map[int]bool{}
    maxLength := 0

    for _, num := range nums {
        checkNum[num] = true
    }

    for _, num := range nums {
        if !checkNum[num-1] {
            count := 0
            for checkNum[num+count] {
                count++
            }
            maxLength = max(maxLength,count)
        }
    }

    return maxLength
}
