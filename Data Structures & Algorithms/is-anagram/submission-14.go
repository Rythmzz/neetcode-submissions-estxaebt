func isAnagram(s string, t string) bool {
	if len(t) != len(s) {
		return false
	}

	seen := make(map[byte]int)

	for i:=0; i< len(t);i++ {
		seen[t[i]]++
	}

	for i:= 0; i< len(s);i++ {
		seen[s[i]]--
	}

	for _ , count := range seen {
		if count != 0 {
			return false
		}
	}

	return true
}
