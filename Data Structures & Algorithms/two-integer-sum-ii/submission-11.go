func twoSum(numbers []int, target int) []int {
    left,right := 0, len(numbers)-1

    for left < right {
        sum := numbers[left] + numbers[right]

        if left < right && sum < target {
            left++
        } else if left < right && sum > target {
            right--
        } else  {
            return []int{left+1,right+1}
        }


    }
    
    return nil
}
