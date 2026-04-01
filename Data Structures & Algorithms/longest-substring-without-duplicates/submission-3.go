func lengthOfLongestSubstring(s string) int {
	maxLength,l,r := 0,0,0
	lastSeen := map[byte]int{}

	for r < len(s) {
		ch := s[r]

		if idx, exists := lastSeen[ch];exists {
			if idx >= l {
				l = idx +1
			}
		}

		lastSeen[ch] = r
		maxLength = max(maxLength,r-l+1)
		r++
	}

	return maxLength
}
