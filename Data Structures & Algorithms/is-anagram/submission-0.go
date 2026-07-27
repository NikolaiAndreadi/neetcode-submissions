func isAnagram(s string, t string) bool {
    fingerprint := make(map[rune]int)
    for _, c := range s {
        fingerprint[c] += 1
    }
    for _, c := range t {
        fingerprint[c] -= 1
    }
    for _, v := range fingerprint {
        if v != 0 {
            return false
        }
    }
    return true
}
