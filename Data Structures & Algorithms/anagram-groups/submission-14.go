func groupAnagrams(strs []string) [][]string {
	frequency := map[[26]int][]string{}

	for _, str := range strs {
		var freq [26]int
		for _,ch := range str {
			freq[ch-'a']++
		}
		frequency[freq] = append(frequency[freq],str)
	}

	result := [][]string{}

	for _, s := range frequency {
		result = append(result,s)
	}

	return result


}
