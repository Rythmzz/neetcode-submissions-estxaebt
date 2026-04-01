func twoSum(nums []int, target int) []int {
    index := map[int]int{}

    for i:= 0 ;i < len(nums);i++ {
        index[nums[i]] = i+1
    }

    for i:=0 ;i < len(nums);i++{
        num := target - nums[i]

        if index[num] != 0 {
            if i != index[num]-1 {
               return []int{i,index[num]-1}
            }
        }
    }

    return nil
}
