func isPalindrome(s string) bool {
	l,r := 0, len(s)-1

	for l < r {
		for l < r && !isAc(s[l]) {
			l++
		}

		for l < r && !isAc(s[r]) {
			r--
		}

		if cvLower(s[l]) != cvLower(s[r]) {
			return false
		}

		l++
		r--
	}

	return true
}

func isAc(c byte) bool {
	if c >= 'a' && c <= 'z' {
		return true
	}

	if c >= 'A' && c <= 'Z' {
		return true
	}

	if c >= '0' && c <= '9' {
		return true
	}

	return false
}

func cvLower(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		return c+32
	}

	return c
}
