func lengthOfLongestSubstring(s string) int {
	maxLength := 0
	lastSeen := map[byte]int{}

	left := 0
	for right:=0;right < len(s);right++ {
		ch := s[right]
		if idx, exists := lastSeen[ch]; exists && idx >= left {
			left = idx + 1
		}

		lastSeen[ch] = right
		maxLength = max(maxLength,right-left+1)
	}

	return maxLength
}
