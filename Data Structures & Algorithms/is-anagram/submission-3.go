func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	checkCharS := map[byte]int{}
	checkCharT := map[byte]int{}

	for i := 0; i < len(s); i++ {
		checkCharS[s[i]]++
		checkCharT[t[i]]++
	}

	for i := 0; i < len(s); i++ {
		if checkCharS[s[i]] != checkCharT[s[i]] {
			return false
		}
	}
	
	return true
}
