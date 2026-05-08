func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	} 

	var a1 [26]int
	var a2 [26]int

	for i:= 0; i < len(s1);i++ {
		a1[s1[i]-'a']++
		a2[s2[i]-'a']++
	}

	if a1 == a2 {
		return true
	}

	for i:= len(s1);i< len(s2);i++ {
		a2[s2[i-len(s1)]-'a']--
		a2[s2[i]-'a']++

		if a1 == a2 {
			return true
		}
	}

	return false
}
