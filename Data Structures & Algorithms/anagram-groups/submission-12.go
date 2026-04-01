func groupAnagrams(strs []string) [][]string {
	frequency := make(map[[26]int][]string)

	for _, str:= range strs {
		var freq [26]int
		for _, ch:= range str {
			freq[ch-'a']++
		}

		frequency[freq] = append(frequency[freq],str)
	}

	var result [][]string

	for _, freq := range frequency {
		result = append(result,freq)
	}

	return result
}
