func characterReplacement(s string, k int) int {
	freq := map[byte]int{}
	freqMax := 0
	maxLength := 0

	// for i:= 0 ; i< len(s);i++{
	// 	freq[s[i]]++
	// }

	l := 0
	for r:= 0 ;r< len(s);r++ {
		ch := s[r]

		freq[ch]++
		freqMax = max(freqMax,freq[ch])
		windowLen := r - l + 1

		if  windowLen - freqMax > k {
			freq[s[l]]--
			l++
		}

		maxLength = max(maxLength,r-l+1)

	}

	return maxLength
}
