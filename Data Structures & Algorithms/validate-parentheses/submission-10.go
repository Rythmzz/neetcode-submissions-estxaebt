func isValid(s string) bool {
    if len(s) <= 1 {
        return false
    }

    if s[0] == ')' || s[0] == ']' || s[0] == '}' {
        return false
    }
 	stack := []byte{}

	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			stack = append(stack, s[i])
		} else {
            if len(stack) == 0 {
                return false
            }
			top := stack[len(stack)-1]
			if s[i] == '}' || s[i] == ']' {
				if top != s[i]-2 {
					return false
				} else {
					stack = stack[:len(stack)-1]
				}
			} else {
				if top != s[i]-1 {
					return false
				} else {
					stack = stack[:len(stack)-1]
				}
			}
		}
	}

    if len(stack) == 0 {
        return true
    } else {
        return false
    }
}
