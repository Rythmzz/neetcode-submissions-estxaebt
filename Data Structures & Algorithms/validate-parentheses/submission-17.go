func isValid(s string) bool {
    var stack []byte

	if !isBracketOpen(s[0]) {
		return false
	} else {
		stack = append(stack,s[0])
	}

	for i:= 1 ; i < len(s);i++ {
		if isBracketOpen(s[i]){
			stack = append(stack,s[i])
		} else if isBracketClose(s[i]) {
			if len(stack) > 0 {
				if convertBracket(s[i]) != stack[len(stack)-1] {
					return false
				} else  {
					stack = stack[:len(stack)-1]
				}
			} else {
				return false
			}
		}
	}

	if len(stack) == 0 {
		return true
	} 

	return false
}

func isBracketOpen(c byte) bool {
	if c == '(' || c == '[' || c== '{' {
		return true
	}

	return false
}

func isBracketClose(c byte)  bool {
	if c == ')' || c == ']' || c == '}'  {
		return true
	}

	return false
}

func convertBracket(c byte) byte {
	switch c {
		case ')': return '('
		case '}': return  '{'
		case ']': return '['
	}
	return c
}
