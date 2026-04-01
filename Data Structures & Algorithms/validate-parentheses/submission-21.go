func isValid(s string) bool {
    var stack []byte

	for i:= 0 ;i < len(s); i++ {
		switch isBracket(s[i]) {
			case 0: {
				stack = append(stack,s[i])
			}
			case 1: {
				if len(stack) > 0 && convertBracketOpen(s[i]) == stack[len(stack)-1] {
					stack = stack[:len(stack)-1]
				} else {
					return false
				}
			}
			case -1: return false
		}
	}
	if len(stack) > 0 {
		return false
	}
	
	return true
}

func isBracket(c byte) int {
	if c == '(' || c == '{' || c == '[' {
		return 0
	}

	if c == ')' || c == '}' || c == ']' {
		return 1
	}

	return -1
}

func convertBracketOpen(c byte) byte {
	if c == ')' { return '('}
	if c == '}' { return '{'}
	if c == ']' { return '['}
	return c
}
