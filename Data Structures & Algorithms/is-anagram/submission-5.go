func isAnagram(s string, t string) bool {
    count := map[byte]int{}

    for i:= 0; i < len(s); i++ {
        count[s[i]]++
    }

    for i:= 0; i < len(t); i++ {
        count[t[i]]--
    }

    for _, c := range count {
        if c != 0 {
            return false
        }
    }

    return true
}
