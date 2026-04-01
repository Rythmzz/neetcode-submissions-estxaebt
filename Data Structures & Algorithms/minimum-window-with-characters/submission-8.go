func minWindow(s string, t string) string {
    if len(t) > len(s) {
		return ""
	}

	need := map[byte]int{}
	for i:= 0; i < len(t);i++ {
		need[t[i]]++
	}

	freq := map[byte]int{}
	l,r,start,length,count := 0,0,0,len(s)+1,0

	for r < len(s) {
		ch := s[r]
		r++

		if need[ch] > 0 {
			freq[ch]++
			if freq[ch] == need[ch] {
				count++
			}
		}

		for count == len(need) {
			if r-l < length {
				start = l
				length = r-l
			}

			d := s[l]
			l++

			if need[d] > 0 {
				if need[d] == freq[d] {
					count--
				}

				freq[d]--
			}
		}
	}

	if length == len(s) + 1 {
		return ""
	}

	return s[start:start+length]
}
