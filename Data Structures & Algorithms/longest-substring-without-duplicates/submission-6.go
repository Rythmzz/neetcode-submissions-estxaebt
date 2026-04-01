func lengthOfLongestSubstring(s string) int {
	lastSeen := map[byte]int{}
	maxLength := 0

	l,r := 0,0
	for r < len(s){
		if idx, exist := lastSeen[s[r]]; exist{
			if idx >= l {
				l = idx+1
			}
		}

		lastSeen[s[r]] = r
		maxLength = max(maxLength,r-l+1)
		r++
	}
	return maxLength
}
