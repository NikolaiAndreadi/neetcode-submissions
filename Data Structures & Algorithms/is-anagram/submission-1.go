func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    
    fingerprint := make(map[byte]int)
    for i := 0; i < len(s); i++ {
        fingerprint[s[i]]++
    }
    for i := 0; i < len(t); i++ {
        fingerprint[t[i]]--
        if fingerprint[t[i]] < 0 {
            return false
        }
    }
    return true
}
