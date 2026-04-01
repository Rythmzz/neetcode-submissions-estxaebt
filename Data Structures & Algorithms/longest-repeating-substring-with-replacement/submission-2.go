func characterReplacement(s string, k int) int {
	freq := map[byte]int{}
	maxLength := 0
	left := 0
	freqMax := 0

	for right:=0 ;right < len(s);right++ {
		ch := s[right]

		freq[ch]++
		freqMax = max(freqMax,freq[ch])

		windowLen := right - left +1
		if windowLen - freqMax > k {
			freq[s[left]]--
			left++
		}

		maxLength = max(maxLength,right-left+1)
	}

	return maxLength


}
