func isAnagram(s string, t string) bool {
	dup := map[byte]int{}

	for i:= 0; i< len(s);i++ {
		dup[s[i]]++
	}

	for i:= 0; i< len(t);i++ {
		dup[t[i]]--
	}

	for _, ch := range dup {
		fmt.Println(ch)
		if ch != 0 {
			return false
		}
	}

	return true
}
