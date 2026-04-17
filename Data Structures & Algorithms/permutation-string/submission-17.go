func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	var ar1 [26]int
	var ar2 [26]int

	for i:=0 ;i< len(s1);i++ {
		ar1[s1[i]-'a']++
		ar2[s2[i]-'a']++
	}

	if ar1 == ar2 {
		return true
	}

	for i:= len(s1);i < len(s2);i++ {
		ar2[s2[i-len(s1)]-'a']--
		ar2[s2[i]-'a']++

		if ar1 == ar2 {
			return true
		}
	}

	return false

}
