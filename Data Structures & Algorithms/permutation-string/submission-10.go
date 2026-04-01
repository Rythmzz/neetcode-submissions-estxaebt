func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		 return false
	}

	countS1 := [26]int{}
	countS2 := [26]int{}

	for i:= 0; i< len(s1);i++{
		countS1[s1[i] - 'a']++
		countS2[s2[i] - 'a']++
	}

	if countS1 == countS2 {
		return true
	}

	for i:= len(s1);i < len(s2);i++{
		countS2[s2[i - len(s1)] - 'a']--
		countS2[s2[i] - 'a']++

		if countS1 == countS2 {
			return true
		}
	}

	return false
}
