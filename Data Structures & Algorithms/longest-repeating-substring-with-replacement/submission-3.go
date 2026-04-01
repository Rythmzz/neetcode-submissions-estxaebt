func characterReplacement(s string, k int) int {
	freq := make(map[byte]int)
	maxLength := 0
	freqMax := 0

	l :=0
	for r:= 0 ; r < len(s); r++ {
		freq[s[r]]++
		freqMax = max(freqMax,freq[s[r]])

		windowLen := r - l + 1
		if windowLen - freqMax > k {
			freq[s[l]]--
			l++
		}

		maxLength = max(maxLength,r - l + 1)
	}

	return maxLength
}
