func minWindow(s string, t string) string {
    if len(t) > len(s) {
		return ""
	}

	need := map[byte]int{}
	for i:= 0 ;i< len(t);i++{
		need[t[i]]++
	}

	l,r,start,valid, length := 0,0,0,0, len(s)+1
	window := map[byte]int{}

	for r < len(s){
		ch := s[r]
		r++

		if need[ch] > 0 {
			window[ch]++
			if need[ch] == window[ch] {
				valid++
			}
		}

		for valid == len(need) {
			if (r - l) < length {
				start = l
				length = r - l
			}

			d := s[l]
			l++

			if window[d] > 0 {
				if need[d] == window[d] {
					valid--
				}
				window[d]--
			}
		}
	}

	if length == len(s)+1 {
		return ""
	}

	return s[start:start+length]
}
