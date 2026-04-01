func longestConsecutive(nums []int) int {
    hasSeen := map[int]bool{}
    maxLength := 0

    for _, num:= range nums {
        hasSeen[num] = true
    }

    for _, num:= range nums {
        if !hasSeen[num-1] {
            count := 0

            for hasSeen[num+count]{
                count++
            }

            maxLength = max(maxLength,count)
        }
    }

    return maxLength
}
