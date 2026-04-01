func dailyTemperatures(temperatures []int) []int {
	result := make([]int,len(temperatures))
	stack := []int{}

	for i:= 0 ;i < len(temperatures);i++ {
		for len(stack) > 0 && temperatures[stack[len(stack)-1]] < temperatures[i] {
			idx := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			result[idx] = i - idx
		}

		stack = append(stack,i)
	}

	return result
}
