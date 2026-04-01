func minWindow(s string, t string) string {
	if len(t) > len(s) {
		return ""
	}


    need := map[byte]int{}

	for i:= 0; i< len(t);i++{
		need[t[i]]++
	}

	window := map[byte]int{}
	l,r,start,length,valid := 0,0,0,len(s)+1,0

	for r < len(s) {
		ch := s[r]
		r++

		window[ch]++
		if window[ch] == need[ch] {
			valid++
		}

		for len(need) == valid {
			if r - l < length {
				start = l
				length = r - l
			}

			dh := s[l]
			l++

			if need[dh] > 0 {
				if window[dh] == need[dh] {
					valid--
				}
				window[dh]--
			}
		}
	}

	if length == len(s)+1 {
		return ""
	}

	return s[start:start+length]
}
