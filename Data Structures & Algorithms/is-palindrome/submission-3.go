func isPalindrome(s string) bool {
    left, right := 0, len(s) - 1

    for left < right {
        
        for  left < right && !isAlphaNumeric(s[left]) {
            left++
        }

        for  left < right && !isAlphaNumeric(s[right]) {
            right--
        }


        if convertLowerAlpha(s[left]) != convertLowerAlpha(s[right]) {
            return false
        }

        left++
        right--
    }

    return true
}

func isAlphaNumeric(ch byte) bool {
    if ch >= 'A' && ch <= 'Z' { return true }
    if ch >= 'a' && ch <= 'z' { return true }
    if ch >= '0' && ch <= '9' { return true }
    
    return false
}

func convertLowerAlpha(ch byte) byte{
    if ch >= 'A' && ch <= 'Z' {
        return ch + 32
    }

    return ch
}


