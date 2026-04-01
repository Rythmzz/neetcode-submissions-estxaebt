func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	frequency := make(map[[26]int][]string)
	var freqS1 [26]int
	for _, ch:= range s1 {
		freqS1[ch-'a']++
	}
	frequency[freqS1] = append(frequency[freqS1],s1)

	var result []string

	for i:= 0 ; i < len(s2);i++ {
		length := i + len(s1)
		if length <= len(s2) {
		result = append(result,s2[i:length])
		} else {
			break
		}
	}

	for _, str := range result {
		var freq[26]int

		for _, ch := range str {
			freq[ch-'a']++
		}
	frequency[freq] = append(frequency[freq],str)
	}

	if len(frequency[freqS1]) > 1 {
		return true
	}

	return false

}
