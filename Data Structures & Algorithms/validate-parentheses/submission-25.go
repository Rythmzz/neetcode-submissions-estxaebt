func isValid(s string) bool {
    st := []byte{}

	for i:= 0; i < len(s);i++ {
		if isBracketOpen(s[i]) {
			st = append(st,s[i])
		} else if isBracketClose(s[i]) {
			if len(st) > 0 {
				if st[len(st)-1] == convertBracketOpen(s[i]) {
					st = st[:len(st)-1]
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

	if len(st) != 0 {
		return false
	}

	return true
}

func isBracketOpen(c byte) bool {
	if c == '(' || c == '[' || c == '{' { 
		return true
	}

	return false
}

func isBracketClose(c byte) bool {
	if c == ')' || c == ']' || c == '}' { 
		return true
	}

	return false
}

func convertBracketOpen(c byte) byte {
	switch c {
		case ')' : return '('
		case ']' : return '['
		case '}' : return '{'
	}

	return c
}


