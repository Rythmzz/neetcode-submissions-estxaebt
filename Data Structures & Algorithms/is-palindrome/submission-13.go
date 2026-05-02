func isPalindrome(s string) bool {
	l,r := 0,len(s)-1

	for l < r {
		for l < r && !isAlphanumeric(s[l]) {
			l++
		}

		for l < r && !isAlphanumeric(s[r]) {
			r--
		}

		if convertLC(s[l]) != convertLC(s[r]) {
			return false
		}

		l++
		r--
	}

	return true
}

func isAlphanumeric(ch byte) bool {
	if ch >= 'A' && ch <= 'Z' {return true}

	if ch >= 'a' && ch <= 'z' {return true}

	if ch >= '0' && ch <= '9' {return true}

	return false
}

func convertLC(ch byte) byte{
	if ch >= 'A' && ch <= 'Z' {return ch+32}

	return ch
}
