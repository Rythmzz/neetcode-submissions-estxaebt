func characterReplacement(s string, k int) int {
	freqMax := 0
	freq := map[byte]int{}

	maxLength := 0
	l := 0
	for r:= 0 ;r < len(s);r++{
		ch := s[r]

		freq[ch]++
		freqMax = max(freqMax, freq[ch])

		windowLen := r - l + 1
		if windowLen - freqMax > k {
			freq[s[l]]--
			l++
		}

		maxLength = max(maxLength, r-l+1)
	}
	return maxLength
}
