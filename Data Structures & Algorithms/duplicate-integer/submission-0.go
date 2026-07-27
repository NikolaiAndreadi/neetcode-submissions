func hasDuplicate(nums []int) bool {
    uniq := make(map[int]struct{})
    for _, n := range nums {
        if _, ok := uniq[n]; ok {
            return true
        }
        uniq[n] = struct{}{}
    }
    return false
}
