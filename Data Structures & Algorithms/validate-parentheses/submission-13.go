func isValid(s string) bool {
	if len(s) <= 1 {
		return false
	}
	var stack []byte
	for i:= 0 ;i< len(s);i++{

		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack = append(stack,s[i])
		} else {
			if len(stack) == 0 {
				return false
			}

			if stack[len(stack)-1] == '(' {
				if s[i] != ')' {
					return false
				}
			}

			if stack[len(stack)-1] == '[' {
				if s[i] != ']' {
					return false
				}
			}

			if stack[len(stack)-1] == '{' {
				if s[i] != '}' {
					return false
				}
			}

			stack = stack[:len(stack)-1]
		}
	}

	if len(stack) != 0 {
		return false
	}

	return true
}
