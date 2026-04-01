func lengthOfLongestSubstring(s string) int {
	lastSeen := map[byte]int{}
	maxLength := 0

	l := 0
	for r:= 0; r< len(s);r++ {
		ch := s[r]

		if idx, exists := lastSeen[ch]; exists {
			if idx >= l {
				l = idx + 1
			}
		}

		lastSeen[ch] = r
		maxLength = max(maxLength, r - l +1)
	}

	return maxLength
}
