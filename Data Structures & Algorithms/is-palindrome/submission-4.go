func isPalindrome(s string) bool {
    left,right := 0, len(s)-1

    for left < right {
        for left < right && !isAlphanumeric(s[left]) {
            left++
        }

        for left < right && !isAlphanumeric(s[right]) {
            right--
        }

        
        if convertLowerCase(s[left]) != convertLowerCase(s[right]){
            return false
        }

        left++
        right--
    }

    return true
}

func isAlphanumeric(ch byte) bool {
    if ch >= 'A' && ch <= 'Z' {
        return true
    }

    if ch >= 'a' && ch <= 'z' {
        return true
    }

    if ch >= '0' && ch <= '9' {
        return true
    }

    return false
}

func convertLowerCase(ch byte) byte {
    if ch >= 'A' && ch <= 'Z' {
        return ch+32
    }

    return ch
}
