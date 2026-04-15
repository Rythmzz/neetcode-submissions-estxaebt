func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	dup := map[byte]int{}

	for i:=0;i<len(t);i++ {
		dup[t[i]]++
	}

	for i:=0;i < len(s);i++ {
		dup[s[i]]--
	}


	for _,count := range dup {
		if count > 0 {
			return false
		}
	}

	return true
}
