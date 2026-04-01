func checkInclusion(s1 string, s2 string) bool {
	return checkPermutation(s1, s2)
}

func checkPermutation(s1 string, s2 string) bool {
	// if len(s1) <= 1 {
	// 	return false
	// }
	counter := map[byte]int{}

	for i := 0; i < len(s1); i++ {
		counter[s1[i]]++
	}

	l := 0
	for r := 0; r < len(s2); r++ {
		ch := s2[r]
		if counter[ch]-1 < 0 {
			l = r
		} else {
			length := r - l + 1
			fmt.Printf("l:%d, r:%d\n", l, r)
			if length == len(s1) {
				str := s2[l : r+1]
				fmt.Printf("s1: %v, s[%d:%d]: %v\n", s1, l, r, str)
				if isPermutation(s1, str) {
					return true
				} else {
					l++
				}
			}
		}
	}

	return false
}

func isPermutation(s1 string, s2 string) bool {
	counter := map[byte]int{}

	for i := 0; i < len(s1); i++ {
		counter[s1[i]]++
	}

	for i := 0; i < len(s2); i++ {
		counter[s2[i]]--
	}

	for i := 0; i < len(s1); i++ {
		if counter[s1[i]] != 0 {
			return false
		}
	}

	return true
}