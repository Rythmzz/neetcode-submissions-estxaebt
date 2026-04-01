func twoSum(nums []int, target int) []int {
    numCheck := map[int]bool{}
	numIndex := map[int]int{}

	for idx, num := range nums {
		numCheck[num] = true
		numIndex[num] = idx
	}

	for idx,num := range nums {
		if numCheck[target-num] {
			if idx != numIndex[target-num]{
			return []int{idx,numIndex[target-num]}
			}
		}
	}

	return nil
}
