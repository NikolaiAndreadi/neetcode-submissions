func dailyTemperatures(temperatures []int) []int {
	result := make([]int, len(temperatures))
	for i := range temperatures {
		ltemp := temperatures[i]
		for j:=i+1; j<len(temperatures); j++ {
			if temperatures[j] > ltemp {
				result[i] = j-i
				break
			}
		}
	}
	return result
}
