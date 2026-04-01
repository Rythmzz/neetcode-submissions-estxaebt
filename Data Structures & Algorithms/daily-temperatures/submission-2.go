func dailyTemperatures(temperatures []int) []int {
	var result []int
	n := len(temperatures)
	maxTemp := temperatures[0]

	for _, temp := range  temperatures {
		maxTemp = max(maxTemp,temp)
	}

	for i:=0;i<n;i++ {
		count := 0
		if temperatures[i] != maxTemp {
			for j:=i+1;j<n;j++{
			if j+1 == n && temperatures[i] >= temperatures[j]{
				count = 0
			} else if temperatures[i] < temperatures[j] {
				count++
				break
			} else {
				count++
			}
			}
		}
		
		result = append(result,count)
	}

	return result
}
