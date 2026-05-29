func isAnagram(s string, t string) bool {
	seen := make(map[rune]int)

	for _, n := range t {
		seen[n]++
	}

	for _, n := range s {
		seen[n]--
	}

	for _, n := range seen {
		if n != 0 {
			return false
		}
	}

	return true
}
