func checkInclusion(s1 string, s2 string) bool {

	if len(s1) > len(s2) {
		return false
	}

	counts1 := [26]int{}
	counts2 := [26]int{}

	for i:= 0 ;i < len(s1);i++{
		counts1[s1[i]-'a']++
		counts2[s2[i]-'a']++
	}

	if counts1 == counts2 {
		return true
	}

	for i:=len(s1) ;i< len(s2);i++ {
		counts2[s2[i-len(s1)]-'a']--
		counts2[s2[i]-'a']++

		if counts1 == counts2 {
			return true
		}
	}

	return false
}
