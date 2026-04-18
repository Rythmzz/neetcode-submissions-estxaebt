func isValid(s string) bool {
    stack := []byte{}

	for i:=0 ;i< len(s);i++ {
		if isBracketOpen(s[i]) {
			stack = append(stack,s[i])
		} else if isBracketClose(s[i]) {
			if len(stack) > 0 {
				if stack[len(stack)-1] != convertBracketOpen(s[i]) {
					return false
				} else {
					stack = stack[:len(stack)-1]
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

func isBracketOpen(ch byte) bool {
	if ch == '(' || ch == '[' || ch == '{' {
		return true
	}

	return false
}

func isBracketClose(ch byte) bool {
	if ch == ')' || ch ==']' || ch == '}' {
		return true
	}

	return false
}

func convertBracketOpen(ch byte) byte {
	if ch == ')'{
		return '('
	}

	if ch == ']'{
		return '['
	}

	if ch == '}' {
		return '{'
	}

	return ch
}
