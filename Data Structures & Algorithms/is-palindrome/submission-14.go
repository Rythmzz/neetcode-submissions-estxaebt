func isPalindrome(s string) bool {
	l,r := 0, len(s)-1

	for l < r {
		for l < r && !isAlphanumeric(s[l]) {
			l++
		}

		for l < r && !isAlphanumeric(s[r]) {
			r--
		}

		if lowerCase(s[l]) != lowerCase(s[r]) {
			return false
		}

		l++
		r--
	}

	return true
}

func isAlphanumeric(c byte) bool {
	if c >= 'A' && c <= 'Z' {
		return true
	}

	if c >= 'a' && c <= 'z' {
		return true
	}

	if c >= '0' && c <= '9' {
		return true
	}

	return false
}

func lowerCase(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		return c+32
	}

	return c
}
