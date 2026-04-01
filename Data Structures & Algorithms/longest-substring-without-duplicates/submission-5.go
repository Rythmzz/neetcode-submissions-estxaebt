func lengthOfLongestSubstring(s string) int {
	lastSeen := map[byte]int{}
	maxLength := 0

	l := 0
	for i:= 0 ; i< len(s);i++{
		ch := s[i]
		if idx, exists := lastSeen[ch]; exists {
			if idx >= l {
				l = idx + 1
			}
		}

		lastSeen[ch] = i
		maxLength = max(maxLength, i - l + 1)
	}

	return maxLength
}
