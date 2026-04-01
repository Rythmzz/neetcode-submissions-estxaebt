func isValid(s string) bool {
    var stack []byte

	for i:=0 ; i< len(s);i++{
		if isBracketOpen(s[i]) {
			stack = append(stack,s[i])
		} else if isBracketClose(s[i]) {
			if len(stack) > 0 {
				if stack[len(stack)-1] == '(' {
					if s[i] != ')'{
						return false
					}
				} else if stack[len(stack)-1] == '[' {
					if s[i] != ']'{
						return false
					}
				} else {
					if s[i] != '}'{
						return false
					}
				}
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

func isBracketOpen(ch byte) bool {
	if ch == '(' || ch ==  '{' || ch ==  '[' {
		return true
	}  


	return false
}

func isBracketClose(ch byte) bool {
	if ch == ')' || ch ==  '}' || ch ==  ']' {
		return true
	}  


	return false
}
