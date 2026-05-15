func dailyTemperatures(temperatures []int) []int {
	st := []int{}
	result := make([]int,len(temperatures))

	for i:= 0; i < len(temperatures);i++ {
		for len(st) > 0 && temperatures[i] > temperatures[st[len(st)-1]] {
			idx := st[len(st)-1]
			st = st[:len(st)-1] 
			result[idx] = i - idx
		}
		st = append(st,i)
	}

	return result
}
