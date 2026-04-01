func isValid(s string) bool {
    var stack []byte

    for i:= 0 ;i < len(s);i++ {
        if isBracketOpen(s[i]) {
            stack = append(stack,s[i])
        } else {
            if isBracketClose(s[i]) {
                if len(stack) > 0 {
                    if convertBacketCloseToOpen(s[i]) == stack[len(stack)-1] {
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
    }

    if len(stack) > 0 {
        return false
    }
    return true
}

func isBracketOpen(ch byte) bool {
    if ch == '(' || ch == '{' || ch == '[' {
        return true
    }

    return false
}

func isBracketClose(ch byte) bool {
    if ch == ')' || ch == '}' || ch == ']' {
        return true
    }

    return false
}

func convertBacketCloseToOpen(ch byte) byte {
    switch ch {
        case ')': return '('
        case ']': return '['
        case '}': return '{'
    }
    return ch
}
