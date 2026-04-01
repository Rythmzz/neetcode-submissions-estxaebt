func longestConsecutive(nums []int) int {
    flagNum := map[int]bool{}
    maxLength := 0

    for _, num := range nums {
        flagNum[num] = true
    }
	
    for _, num := range nums {
        if !flagNum[num-1] {
            i := num
			count:= 0
			for flagNum[i] {
				count++
				i++
			}
			maxLength = max(maxLength,count)
        }
    }

	return maxLength

}
