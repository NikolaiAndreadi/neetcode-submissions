func trap(height []int) int {
	// 0 2 2 3 3 3 3 3 3 3
	// 3 3 3 3 3 3 3 3 2 1
	// 0 2 2 3 3 3 3 2 2 1
	l := 0
	r := len(height)-1
	maxL := height[l]
	maxR := height[r]
	result := 0
	for l < r {
		if maxL < maxR {
			l++
			maxL = max(maxL, height[l])
			delta := min(maxL, maxR) - height[l]
			if delta > 0 {
				result += delta
			}
		} else {
			r--
			maxR = max(maxR, height[r])
			delta := min(maxL, maxR) - height[r]
			if delta > 0 {
				result += delta
			}
		}
	}
	return result
}