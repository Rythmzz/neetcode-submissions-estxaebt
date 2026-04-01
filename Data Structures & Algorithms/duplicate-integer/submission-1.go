func hasDuplicate(nums []int) bool {
    exist := map[int]bool{}

    for _, num := range nums {
        if exist[num] {
            return true
        }
        exist[num] = true
    }

    return false
}
