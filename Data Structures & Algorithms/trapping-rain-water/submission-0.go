func trap(height []int) int {
	// 0 2 2 3 3 3 3 3 3 3
	// 3 3 3 3 3 3 3 3 2 1
	// 0 2 2 3 3 3 3 2 2 1
	
	result := 0
	tmp := make([]int, len(height))
	prevMax := 0
	for i, v := range height {
		prevMax = max(prevMax, v)
		tmp[i] = prevMax
	}
	prevMax = 0
	for i := len(height)-1; i > 0; i-- {
		prevMax = max(prevMax, height[i])
		tmp[i] = min(tmp[i], prevMax)
	}
	for i := range height {
		delta := tmp[i] - height[i]
		if delta <= 0 {
			continue
		}
		result += delta
	}
	return result
}
