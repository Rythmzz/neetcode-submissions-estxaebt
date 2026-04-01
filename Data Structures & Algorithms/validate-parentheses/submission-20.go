func isValid(s string) bool {
    var stack []byte

	for i:= 0 ;i< len(s);i++ {
		if isOpenBracket(s[i]) {
			stack = append(stack,s[i])
		} else if isCloseBracket(s[i]) {
			if len(stack) == 0 {
				return false
			}

			c := stack[len(stack)-1]
			if c == convertBracket(s[i]) {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		} else {
			return false
		}
	}

	if len(stack) > 0 {
		return false
	}

	return true
}

func convertBracket(ch byte) byte{
	if ch == ')' {return '('}
	if ch == ']' {return '['}
	if ch == '}' {return '{'}
	return ch
}

func isOpenBracket(ch byte) bool {
	if ch == '(' || ch == '{' || ch == '[' {
		return true
	}

	return false
}

func isCloseBracket(ch byte) bool {
	if ch == ')' || ch == '}' || ch == ']' {
		return true
	}

	return false
}
