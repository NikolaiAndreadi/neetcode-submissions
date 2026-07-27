func groupAnagrams(strs []string) [][]string {
	type fingerprint [32]int
	result := make(map[fingerprint][]string)
	for _, str := range strs {
		var fp fingerprint
		for _, c := range str {
			chNum := int(c) - int('a')
			fp[chNum]++
		}
		subset, ok := result[fp]
		if ok {
			result[fp] = append(subset, str)
		} else {
			result[fp] = []string{str}
		}
	}
	actualResult := make([][]string, len(result))
	i := 0
	for _, v := range result {
		actualResult[i] = v
		i++
	}
	return actualResult
}
