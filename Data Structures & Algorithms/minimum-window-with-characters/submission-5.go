func minWindow(s string, t string) string {
    if len(s) < len(t) {
		return ""
	}

	freq := map[byte]int{}

	for i:= 0;i< len(t);i++{
		freq[t[i]]++
	}

	need := map[byte]int{}
	l,r,start,valid := 0,0,0,0
	length := len(s)+1

	for r < len(s) {
		ch := s[r]
		r++

		if freq[ch] > 0 {
			need[ch]++
			if need[ch] == freq[ch] {
				valid++
			}
		}

		for valid == len(freq) {
			if r - l < length {
				start = l
				length = r - l
			}

			d := s[l]
			l++

			if need[d] > 0 {
				if need[d] == freq[d] {
					valid--
				}
				need[d]--
			}
		}
	}

	if length == len(s)+1 {
		return ""
	}
	return s[start:start+length]
}
