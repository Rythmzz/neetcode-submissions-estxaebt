func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	ls1 := [26]int{}
	ls2 := [26]int{}

	for i:= 0 ; i< len(s1);i++ {
		ls1[s1[i]-'a']++
		ls2[s2[i]-'a']++
	}

	if ls1 == ls2 {
		return true
	}

	for i:= len(s1);i < len(s2);i++{
		ls2[s2[i - len(s1)]-'a']--
		ls2[s2[i]-'a']++

		if ls1 == ls2 {
			return true
		}
	}

	return false
}
