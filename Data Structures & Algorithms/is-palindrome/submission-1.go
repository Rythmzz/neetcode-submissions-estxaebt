func isPalindrome(s string) bool {
    if len(s) == 0 {
        return true
    }
    reg := regexp.MustCompile("[^a-zA-Z0-9]+")
    s = reg.ReplaceAllString(s, "")
    s = strings.ToLower(s)
        
    left,right := 0, len(s)-1

    for left < right {
        // fmt.Println(s[left],s[right])
        if s[left] != s[right]{
            return false
        }
        left++
        right--
    }

    return true
}
