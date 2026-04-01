func isAnagram(s string, t string) bool {
	anagram := map[byte]int{}

	for i:=0 ;i< len(s);i++ {
		anagram[s[i]]++
	}

	for i:=0 ;i< len(t);i++ {
		if anagram[t[i]] > 0 {
			anagram[t[i]]--
		} else {
			return false
		}
	}

	for i:=0 ;i< len(s);i++ {
		if anagram[s[i]] != 0 {
			return false
		}
	}

	return true
	
}
