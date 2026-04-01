func characterReplacement(s string, k int) int {
	freq := map[byte]int{}
	freqMax := 0
	maxLength := 0

	l := 0
	for i:= 0 ; i< len(s);i++ {
		ch := s[i]
		freq[ch]++

		freqMax = max(freqMax,freq[ch])
		windowLen := i - l + 1
		if windowLen - freqMax > k {
			freq[s[l]]--
			l++
		}

		maxLength = max(maxLength,i-l+1)
	}

	return maxLength
}
