func lengthOfLongestSubstring(s string) int {
	seen := make(map[byte]int)
	maxLength := 0
	l := 0
	for r := 0;r < len(s);r++ {
		ch := s[r]
		if idx, exists := seen[ch]; exists {
			if idx >= l {
				l = idx + 1
			}
		}

		seen[ch] = r
		maxLength = max(maxLength,r-l+1)
		
	}

	return maxLength
}
