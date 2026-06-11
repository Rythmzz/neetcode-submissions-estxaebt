func characterReplacement(s string, k int) int {
	freq := map[byte]int{}
	maxFreq := 0
	maxLength := 0

	l := 0
	for r := 0; r < len(s); r++ {
		ch := s[r]
		freq[ch]++
		maxFreq = max(maxFreq, freq[ch])
		windowLen := r - l + 1
		if windowLen-maxFreq > k {
			freq[s[l]]--
			l++
		}

		maxLength = max(maxLength, r-l+1)
	}
	
	return maxLength
}
