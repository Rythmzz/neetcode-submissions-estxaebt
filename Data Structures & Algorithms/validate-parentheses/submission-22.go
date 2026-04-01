func isValid(s string) bool {
    var stack []byte

	for i:= 0; i< len(s);i++ {
		c := s[i]
		if c == '(' || c == '{' || c == '[' {
			stack = append(stack,c)
		} else if c == ')' || c == '}' || c == ']' {
			if len(stack) > 0 {
				switch c {
					case ')': c = '('
					case '}': c = '{'
					case ']': c = '['
				}

				if c != stack[len(stack)-1] {
					return false
				}
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		} else {
			return false
		}
	}

	if len(stack) == 0 {
		return true
	}

	return false
}
