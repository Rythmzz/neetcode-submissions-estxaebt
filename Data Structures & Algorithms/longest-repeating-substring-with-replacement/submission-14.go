func characterReplacement(s string, k int) int {
	maxLength := 0
	maxFreq := 0
	freq := map[byte]int{}

	l := 0
	for r:=0;r < len(s);r++ {
		ch := s[r]
		freq[ch]++

		maxFreq = max(maxFreq,freq[ch])
		windowLen := r - l + 1

		if windowLen - maxFreq > k {
			dh := s[l]
			freq[dh]--
			l++
		}

		maxLength = max(maxLength,r-l+1)
	}

	return maxLength
}
