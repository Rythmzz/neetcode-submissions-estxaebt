func characterReplacement(s string, k int) int {
	freq := map[byte]int{}
	maxFreq := 0
	maxLength := 0

	l,r := 0,0

	for r < len(s) {
		ch := s[r]
		
		freq[ch]++
		maxFreq = max(maxFreq,freq[ch])
		windowLen := r - l + 1
		if windowLen - maxFreq > k {
			freq[s[l]]--
			l++
		}

		maxLength = max(maxLength,r-l+1)
		r++
	}

	return maxLength

}
