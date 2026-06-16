func isValid(s string) bool {
    stack := []byte{}

	for i:=0; i< len(s);i++ {
		c := s[i]
		if isOp(c) {
			stack = append(stack,c)
		} else if isCl(c) {
			if len(stack) > 0 {
				if cCl(stack[len(stack)-1]) == c {
					stack = stack[:len(stack)-1]
				} else {
					return false
				}
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

func isOp(c byte) bool {
	if c == '(' || c == '[' || c == '{' {
		return true
	}

	return false
}

func isCl(c byte) bool {
	if c == ')' || c == ']' || c == '}' {
		return true
	}

	return false
}

func cCl(c byte) byte {
	switch c {
		case '(':return ')'
		case '[':return ']'
		case '{':return '}'
	}

	return c
}
