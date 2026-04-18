func dailyTemperatures(temperatures []int) []int {
	stack := []int{}
	result := make([]int,len(temperatures))

	for i:= 0; i< len(temperatures);i++ {
		for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			idx := stack[len(stack)-1]
			result[idx] = i - idx
			stack = stack[:len(stack)-1]
		}
		stack = append(stack,i)
	}

	return result
}
