func isValid(s string) bool {
	stack := []byte{}
    for i:= 0 ;i< len(s);i++ {
		ch := s[i]
		if isBracketOpen(ch) {
			stack = append(stack,ch)
		} else if isBracketClose(ch) {
			if len(stack) == 0 {
				return false
			}

			if convertBracket(ch) == stack[len(stack)-1] {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		} else {
			return false
		}
	}

	return len(stack) == 0
}

func isBracketOpen(ch byte) bool {
	if ch == '(' || ch == '[' || ch == '{' {
		return true
	} 

	return false
}

func isBracketClose(ch byte) bool {
	if ch == ')' || ch == ']' || ch == '}' {
		return true
	} 

	return false
}

func convertBracket(ch byte) byte {
	if ch == ')' {return '('}
	if ch == '}' {return '{'}
	if ch == ']' {return '['}

	return ch
}




