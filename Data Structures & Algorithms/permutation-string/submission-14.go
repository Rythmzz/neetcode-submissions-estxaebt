func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2){
		return false
	}

	sl1 := [26]int{}
	sl2 := [26]int{}

	for i:= 0;i < len(s1);i++ {
		sl1[s1[i]-'a']++
		sl2[s2[i]-'a']++
	}

	if sl1 == sl2 {
		return true
	}

	for i:= len(s1);i < len(s2);i++ {
		sl2[s2[i-len(s1)]-'a']--
		sl2[s2[i]-'a']++

		if sl1 == sl2 {
			return true
		}
	}

	return false


}
